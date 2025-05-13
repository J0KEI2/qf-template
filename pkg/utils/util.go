package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/valyala/fastjson"
)

var (
	// Session storage
	SessStore *session.Store
	// parser pool
	JsonParserPool *fastjson.ParserPool
)

func init() {
	if JsonParserPool == nil {
		JsonParserPool = new(fastjson.ParserPool)
	}
}

func EncryptSSN(ssn *string) *string {
	if ssn != nil {
		encryptSSN := base64.StdEncoding.EncodeToString([]byte(GetStringFromPointer(ssn)))
		return &encryptSSN
	}
	return nil
}

func CreateEnum(enumName string, enumValue ...string) string {
	return `
	DO $$ BEGIN
			CREATE TYPE ` + enumName + ` AS ENUM('` + strings.Join(enumValue, "', '") + `');
	EXCEPTION
			WHEN duplicate_object THEN null;
	END $$;`
}

func StringifyStringArray(stringArr []string) string {
	outByte, err := json.Marshal(stringArr)
	if err != nil {
		return ""
	}
	return string(outByte)
}

func ParseStringArray(str string) []string {
	out := []string{}
	err := json.Unmarshal([]byte(str), &out)
	if err != nil {
		return []string{}
	}
	return out
}

func IsInRangeUINT(desireInt, min, max uint) bool {
	return desireInt >= min && desireInt <= max
}
