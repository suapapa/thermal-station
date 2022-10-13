# Thermal Station

This project combines following projects:
- [img-receipt](https://github.com/suapapa/img-receipt) : Print image on receipt printer
- [pr_label](https://github.com/suapapa/pr_label) : Print address label on ql800, brother's label printer (uing brother-ql)
- [gb-noti](https://github.com/suapapa/gb-noti) : Print guestboot on receipt printer

## API

### 주문 목록
- `/v1/["receipt"|"label"]/ord`

### 주소 출력
- `/v1/["receipt"|"label"]/addr`

### 이미지 출력
- `/v1/["receipt"|"label"]/img`
```bash
curl -F "img=@./_img/Lenna.png" http://opi-hangulclock.local:8080/v1/receipt/img
```

### QR코드 출력
- `/v1/["receipt"|"label"]/qr`
```bash
curl -X POST -d `{"content": "https://homin.dev"}` http://opi-hangulclock.local:8080/v1/label/qr
```


## MQTT

[방명록 출력 시스템](https://homin.dev/blog/post/20220910_live_print_guestbook_with_mqtt/)에서
영수증 프린터로 방명록을 출력하는 프로그램.


## Tests

```bash
curl -X POST \
-d '{"ID":"1234567890","from":{"line1":"경기 성남시 분당구 판교역로 235 (에이치 스퀘어 엔동)","line2":"7층","name":"카카오 엔터프라이즈","phone_number":"010-1234-5678"},"to":{"line1":"경기도 성남시 분당구 판교역로 166","name":"판교 아지트","phone_number":"010-1234-5678"}}' \
http://rpi-airplay.local:8080/v1/label/order
```