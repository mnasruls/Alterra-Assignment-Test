# Alterra-Assignment-Test

Simple API project using echo framework

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/76ac347865d54ef7a929f9b17e98a34d)](https://www.codacy.com/gh/mnasruls/Alterra-Assignment-Test/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mnasruls/Alterra-Assignment-Test&amp;utm_campaign=Badge_Grade)


## About The Project

Project ini adalah sebuah aplikasi untuk manajemen user dimana terdapat dua entitas yaitu user dan admin. Adapun admin bisa melakukan manajemen user berupa create, update dan delete akun user. User dan admin dapat melihat profile diri, melihat profile user lain dan semua data user.

Berikut fitur yang terdapat dalam Bringeee :

- Fitur Admin
    - Login
    - Create user
    - Edit user
    - Delete user
    - Get profile
    - Get profile user lain
    - Get all user
- Fitur User
    - login
    - Get profile
    - Get profile user lain
    - Get all user

## Tech Stack

![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![Echo](https://img.shields.io/badge/-Echo-05122A?style=flat&logo=go)&nbsp;
![Gorm](https://img.shields.io/badge/-Gorm-05122A?style=flat&logo=go)&nbsp;
![Codacy](https://img.shields.io/badge/-Codacy-05122A?style=flat&logo=codacy)&nbsp;
![JWT](https://img.shields.io/badge/-JWT-05122A?style=flat&logo=jwt)&nbsp;
![Swagger](https://img.shields.io/badge/-Swagger-05122A?style=flat&logo=swagger)&nbsp;

## Link Documentation (Open API)
https://app.swaggerhub.com/apis-docs/nasrulmuhammad63/user-management/v1.0.0
<img src="images/swagger.png">
<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

Kloning project
```bash
  https://github.com/mnasruls/Alterra-Assignment-Test.git
```
Masuk ke direktori project

```bash
  $ cd ~/[nama project]
```
Buat sebuah file dengan nama di dalam folder root project `.env` dengan format `.env.example` yang ada di repository. Sesuaikan configurasi di komputer lokal.

Jalankan aplikasi 

```bash
  $ go run .
```