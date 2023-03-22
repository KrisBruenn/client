package client

import "errors"
import "fmt"

/*
boolean
character
numeric
temporal
*/

func TypeMenu() {
    fmt.Println("*******************")
    fmt.Println("* 1) Boolean      *")
    fmt.Println("* 2) Character    *")
    fmt.Println("* 3) Numeric      *")
    fmt.Println("* 4) Temporal     *")
    fmt.Println("*                 *")
    fmt.Println("* x) Exit         *")
    fmt.Println("*******************")
}

func CharMenu() {
    fmt.Println("*******************")
    fmt.Println("* 1) Fixed char   *")
    fmt.Println("* 2) Variable char*")
    fmt.Println("* 3) Unlimited    *")
    fmt.Println("*                 *")
    fmt.Println("* x) Exit         *")
    fmt.Println("*******************")
}

func numMenu() {

    fmt.Println("*******************")
    fmt.Println("* 1) Integer      *")
    fmt.Println("* 2) Precision    *")
    fmt.Println("* 3) Real         *")
    fmt.Println("* 4) Serial       *")
    fmt.Println("*                 *")
    fmt.Println("* x) Exit         *")
    fmt.Println("*******************")
}

func GetColType() (string, error) {
    var entry, input string
    TypeMenu()
    fmt.Print("Enter type: ")
    fmt.Scanln(&input)
    entry = input[:1]
    switch entry {
        case "x":
            return "", errors.New("Must specify column type!")
        case "1":
            return "BOOL", nil
        case "2":
            var cEntry, cInput, cResult, cSize string
            CharMenu()
            fmt.Print("Enter character type: ")
            fmt.Scanln(&cInput)
            cEntry = cInput[:1]
            switch cEntry {
                case 1:
                    fmt.Print("Enter size: ")
                    fmt.Scanln(&cSize)
                    cResult = fmt.Sprintf("CHAR(%s)", cSize)
                    return cResult, nil
                case 2:
                    fmt.Print("Enter size: ")
                    fmt.Scanln(&cSize)
                    cResult = fmt.Sprintf("VARCHAR(%s)", cSize)
                    return cResult, nil
                case 3:
                    return "TEXT", nil
                case "x":
                    return "", errors.New("Must specify character type!")
        case "3":
            var nEntry, nInput, nResult, nSize string
            NumMenu()
            fmt.Print("Enter numeric type: ")
            fmt.Scanln(&nInput)
            nEntry = nInput[:1]
            switch nEntry {

                case 1:
                    fmt.Print("Enter size [2, 4, 8]: ")
                    fmt.Scanln(&nSize)
                    switch nSize[:1] {
                        case "2":
                            return "SMALLINT", nil
                        case "4":
                            return "INT", nil
                        case "8":
                            return "BIGINT", nil
                        default:
                            return "", errors.New("Invalid integer size!")
                    }
                case 2:
                    var precision, scale string
                    fmt.Print("Enter precision: ")
                    fmt.Scanln(&nSize)
                    precision = nSize
                    fmt.Print("Enter scale: ")
                    fmt.Scanln(&nSize)
                    scale = nSize
                    nResult = fmt.Sprintf("NUMERIC(%s, %s)", precision, scale)
                    return nResult, nil
                case 3:
                    fmt.Print("Enter size [4, 8]: ")
                    fmt.Scanln(&nSize)
                    switch nSize[:1] {
                        case "4":
                            return "REAL", nil
                        case "8":
                            return "DOUBLE PRECISION", nil
                        default:
                            return "", errors.New("Invalid real size!")
                case 4:
                    fmt.Print("Enter size [2, 4, 8]: ")
                    fmt.Scanln(&nSize)
                    switch nSize[:1] {
                        case "2":
                            return "SMALLSERIAL", nil
                        case "4":
                            return "SERIAL", nil
                        case "8":
                            return "BIGSERIAL", nil
                        default:
                            return "", errors.New("Invalid serial size!")
                    }
                case "x":
                    return "", errors.New("Must specify numeric type!")
            }
            
        case "4":
    }
}