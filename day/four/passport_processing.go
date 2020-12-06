package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
  "strconv"
  "regexp"
)

type Passport struct {
  byr string
  iyr string
  eyr string
  hgt string
  hcl string
  ecl string
  pid string
  cid string
}

const FILE_NAME = "./input.txt"

func (p *Passport) AddField(k string, v string) {
  switch k {
    case "byr":
      p.byr = v
    case "iyr":
      p.iyr = v
    case "eyr":
      p.eyr = v
    case "hgt":
      p.hgt = v
    case "hcl":
      p.hcl = v
    case "ecl":
      p.ecl = v
    case "pid":
      p.pid = v
    case "cid":
      p.cid = v
    default:
      fmt.Println("Invalid field:", k)
  }
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func (p *Passport) IsValid() bool {
  if p.byr == "" {
    return false
  } else {
    birth_year, err := strconv.Atoi(p.byr)
    check(err)

    is_valid_birth_year := (birth_year >= 1920) && (birth_year <= 2002)

    if !is_valid_birth_year {
      // fmt.Println("Birth year", p.byr, "is invalid")
      return false
    }
  }

  if p.iyr == "" {
    return false
  } else {
    issue_year, err := strconv.Atoi(p.iyr)
    check(err)
    is_valid_issue_year := issue_year >= 2010 && issue_year <= 2020

    if !is_valid_issue_year {
      // fmt.Println("Issue year", p.iyr, "is invalid")
      return false
    }
  }

  if p.eyr == "" {
    return false
  } else {
    expire_year, err := strconv.Atoi(p.eyr)
    check(err)
    is_valid_expire_year := (expire_year >= 2020 && expire_year <= 2030)

    if !is_valid_expire_year {
      // fmt.Println("Expire year", p.eyr, "is invalid")
      return false
    }
  }

  if p.hgt == "" {
    return false
  } else {
    is_valid_hgt := true

    if strings.HasSuffix(p.hgt, "cm") {
      cms, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm"))
      if (err != nil) {
        is_valid_hgt = false
      }

      if (cms < 150) || (cms > 193) {
        is_valid_hgt = false
      }
    } else if strings.HasSuffix(p.hgt, "in") {
      inches, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in"))
      if (err != nil) {
        is_valid_hgt = false
      }

      if (inches < 59) || (inches > 76) {
        is_valid_hgt = false
      }
    } else {
      is_valid_hgt = false
    }

    if !is_valid_hgt {
      fmt.Println("Height", p.hgt, "is invalid")
      return false
    }
  }

  if p.hcl == "" {
    return false
  } else {
    is_valid_hair_color, err := regexp.MatchString(`^#[a-z\d]{6}$`, p.hcl)
    check(err)

    if !is_valid_hair_color {
      fmt.Println("Hair color", p.hcl, "is invalid")
      return false
    }
  }

  if p.ecl == "" {
    return false
  } else {
    valid_eye_colors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

    is_valid_eye_color := false

    for _, valid_ecl := range(valid_eye_colors) {
      if p.ecl == valid_ecl {
        is_valid_eye_color = true
      }
    }

    if !is_valid_eye_color {
      fmt.Println("Eye color", p.ecl, "is invalid")
      return false
    }
  }

  if p.pid == "" {
    return false
  } else {
    is_valid_passport_id, err := regexp.MatchString(`^[\d]{9}$`, p.pid)
    check(err)

    if !is_valid_passport_id {
      return false
    }
  }

  return true
}


func parse_input(file_name string) *[]*Passport {
  f, err := os.Open(file_name)

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  passports := make([]*Passport, 0)
  curr_passport := new(Passport)

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    row_text := scanner.Text()

    if row_text == "" {
      passports = append(passports, curr_passport)
      curr_passport = new(Passport)
      continue
    }

    fmt.Println(row_text)

    fields := strings.Split(row_text, " ")

    for _, kv_string := range(fields) {
      kv := strings.Split(kv_string, ":")

      curr_passport.AddField(kv[0], kv[1])
    }
  }

  passports = append(passports, curr_passport)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return &passports
}

func count_valid_passports(passports *[]*Passport) int {
  valid_passports := 0

  for _, p := range(*passports) {

    if (p.IsValid()) {
      fmt.Println(p)
      valid_passports++
    }
  }

  return valid_passports
}

func main() {
  passports := parse_input(FILE_NAME)

  fmt.Println("Found", len(*passports), "Passports")

  valid_passports := count_valid_passports(passports)

  fmt.Println("Found", valid_passports, "valid passports")
}
