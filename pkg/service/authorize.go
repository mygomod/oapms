package service

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	mt "math/rand"
	"oapms/pkg/mus"
	"strconv"
	"strings"
)

type authorize struct {
	secret             string
	saltSize           int
	delimiter          string
	stretchingPassword int
}

func InitAuthorize() *authorize {
	obj := &authorize{
		secret:             viper.GetString("oauth.secret"),
		saltSize:           viper.GetInt("oauth.salt"),
		delimiter:          viper.GetString("oauth.delimiter"),
		stretchingPassword: viper.GetInt("oauth.strectchingPassword"),
	}
	return obj
}

func (p *authorize) Hash(pass string) (string, error) {
	salt_secret, err := salt_secret()
	if err != nil {
		return "", err
	}

	salt, err := p.salt(p.secret + salt_secret)
	if err != nil {
		return "", err
	}

	iteration := randInt(1, 20)
	hash, err := p.hash(pass, salt_secret, salt, int64(iteration))
	if err != nil {
		return "", err
	}
	iteration_string := strconv.Itoa(iteration)
	password := p.hashJoin(salt_secret, iteration_string, hash, salt)
	return password, nil

}

//校验密码是否有效
func (p *authorize) Verify(hashing string, pass string) error {
	data := p.trim_salt_hash(hashing)

	iteration, _ := strconv.ParseInt(data["iteration_string"], 10, 64)

	has, err := p.hash(pass, data["salt_secret"], data["salt"], int64(iteration))
	if err != nil {
		return err
	}

	hashJoin := p.hashJoin(data["salt_secret"], data["iteration_string"], has, data["salt"])
	if hashJoin == hashing {
		return nil
	}
	mus.Logger.Debug("util verify error", zap.String("pwd", hashing), zap.String("repwd", hashJoin))
	return errors.New("not equal")

}

func (p *authorize) hash(pass string, salt_secret string, salt string, iteration int64) (string, error) {
	var pass_salt string = salt_secret + pass + salt + salt_secret + pass + salt + pass + pass + salt
	var i int

	hash_pass := p.secret
	hash_start := sha512.New()
	hash_center := sha256.New()
	hash_output := sha256.New224()

	i = 0
	for i <= p.stretchingPassword {
		i = i + 1
		hash_start.Write([]byte(pass_salt + hash_pass))
		hash_pass = hex.EncodeToString(hash_start.Sum(nil))
	}

	i = 0
	for int64(i) <= iteration {
		i = i + 1
		hash_pass = hash_pass + hash_pass
	}

	i = 0
	for i <= p.stretchingPassword {
		i = i + 1
		hash_center.Write([]byte(hash_pass + salt_secret))
		hash_pass = hex.EncodeToString(hash_center.Sum(nil))
	}
	hash_output.Write([]byte(hash_pass + p.secret))
	hash_pass = hex.EncodeToString(hash_output.Sum(nil))

	return hash_pass, nil
}

func (p *authorize) trim_salt_hash(hash string) map[string]string {
	str := strings.Split(hash, p.delimiter)

	return map[string]string{
		"salt_secret":      str[0],
		"iteration_string": str[1],
		"hash":             str[2],
		"salt":             str[3],
	}
}

func (p *authorize) hashJoin(saltSecret, iteration, hash, salt string) string {
	return strings.Join([]string{saltSecret, iteration, hash, salt}, p.delimiter)
}

func (p *authorize) salt(secret string) (string, error) {
	buf := make([]byte, p.saltSize, p.saltSize+md5.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(buf)
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(buf)), nil
}

func salt_secret() (string, error) {
	rb := make([]byte, randInt(10, 100))
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}

func randInt(min int, max int) int {
	return min + mt.Intn(max-min)
}
