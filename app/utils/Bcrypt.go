package utils

import "golang.org/x/crypto/bcrypt"

// untuk menghasilkan hash password pass buat akun dari string password
func HashGenerator(str string)(string,error){
	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	
	return string(hashedString), nil
}

//untuk compare password pas login sama aoa gak sama hashed password di DB
func HashComparator(hashedByte []byte, passwordByte []byte,)error{
	err := bcrypt.CompareHashAndPassword(hashedByte,passwordByte)
	if err != nil {
		return err
		
	}
	return nil
}
