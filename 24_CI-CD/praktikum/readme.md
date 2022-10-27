# Praktikum CI/CD

## Intro
Pada praktikum kali ini, saya mendeploy aplikasi `https://github.com/goFrendiAsgard/alta-batch-3-ec2` pada droplets/instances yang saya di Provider Digital Ocean di IP `128.199.235.210`.
Hasil deployment dapat diakses di [sini](http://128.199.235.210:8080/).

## Proses Instalasi
Setup Firewall
![img_2.png](img_2.png)
Download requirement yang dibutuhkan & build docker
![img.png](img.png)
Docker sudah berjalan
![img_1.png](img_1.png)

## Proses Deployment
Github Actions sudah berjalan ketika ada trigger 
![img_3.png](img_3.png)
Github Actions gagal karena ada kesalahan autentikasi
![img_4.png](img_4.png)
Github Actions berhasil dan sudah di deploy
![img_5.png](img_5.png)

## Hasil Deployment
Yang awalnya memiliki output `Hello World!` menjadi `hi, razanfawwaz here! tugas ci/cd`
![img_6.png](img_6.png)