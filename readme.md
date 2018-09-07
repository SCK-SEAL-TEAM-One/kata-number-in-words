# Kata: Number in word
It occurs now and then in real life that people want to write about money, especially about a certain amount of money. If it comes to cheques or contracts for example some nations have laws that state that you should write out the amount in words additionally to the amount in numbers to avoid fraud and mistakes
100000.00 => หนึ่งแสนบาทถ้วน
400000.00 => สี่แสนบาทถ้วน
 50000.00 => ห้าแสนบาทถ้วน
  5000.00 => ห้าพันบาทถ้วน
   250.50 => สองร้อยห้าสิบบาทห้าสิบสตางค์
2000000.00 => สองล้านบาทถ้วน

# Function And Parameter Naming Conventions

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
api.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_RequestToOtherWebservice_Input_ApplicationForm_Should_Be_ApplicationResult
```

## Variable Name
- ชื่อตัวแปรเป็น **camelCase** เช่น
```
policy, applicationForms, applicationResult, policyResult
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
insurers
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
ResultDay, ResultMonth
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวพิมพ์เล็กก่อน เว้นแต่เมื่อมีการใช้ข้าม package ถึงจะใช้ Capital Case เช่น
```
Hour, Minute, url
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน

## คำสั่ง Run Test
### ค่าสั่ง Run Acceptance Test (Robot)

```
robot duration.robot
```

### คำสั่ง Run Acceptance Test (API)
```
newman run filename
```

## Additional
[See more git and go command](https://github.com/ImKK-000/git-and-go-step)