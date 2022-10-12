# Thermal Station

This is a combine of following projects:
- https://github.com/suapapa/img_receipt : Print image on receipt printer
- https://github.com/suapapa/pr_label : Print address label on ql800, brother's label printer (uing brother-ql)
- https://github.com/suapapa/gb_noti : Print guestboot on receipt printer

## API

## MQTT


## Tests

```bash
curl -X POST \
-d '{"ID":"1234567890","from":{"line1":"경기 성남시 분당구 판교역로 235 (에이치 스퀘어 엔동)","line2":"7층","name":"카카오 엔터프라이즈","phone_number":"010-1234-5678"},"to":{"line1":"경기도 성남시 분당구 판교역로 166","name":"판교 아지트","phone_number":"010-1234-5678"}}' \
http://rpi-airplay.local:8080/v1/label/order
```