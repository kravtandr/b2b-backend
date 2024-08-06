from yookassa import Payment
from yookassa import Configuration
import uuid

from myLogger import logger_decorator

UKASSA_SECRET_KEY = "test_on5XfensBJfl8nY63uzUp3CBTW5YE0w6SVhr4VxAdH4"
Configuration.account_id = "415910" 
Configuration.secret_key = UKASSA_SECRET_KEY

@logger_decorator
def createPayment(value):
    idempotence_key = str(uuid.uuid4())
    value = str(value)+".00"
    payment = Payment.create({
        "amount": {
          "value": value, #"2.00"
          "currency": "RUB"
        },
        "confirmation": {
          "type": "redirect",
          "return_url": "https://bi-tu-bi.ru/profile"
        },
        "payment_method_data": {
          "type": "sbp",
        },
        "description": "Битуби, пополнение баланса: " + value + " рублей",
        "capture": True
    }, idempotence_key)
    return payment

@logger_decorator
def getPaymentInfo(payment_id):
    payment_info = Payment.find_one(payment_id)
    return payment_info


@logger_decorator
def confirmPayment(payment_id, value): 
    idempotence_key = str(uuid.uuid4())

    payment_info = Payment.find_one(payment_id)
    if payment_info.status == "waiting_for_capture" and payment_info.paid == True:
        print("Оплата подтверждена")
        response = Payment.capture(
        payment_id,
        {
            "amount": {
            "value": value, # "2.00",
            "currency": "RUB"
            }
        },
        idempotence_key
        )
        return True
    elif payment_info.status == "succeeded" and payment_info.paid == True:
        print("Оплата подтверждена")
        return True
    else:
        print(payment_info.paid, payment_info.status)
        print("Не оплачено")
        return False

def cancelPayment(payment_id):
  idempotence_key = str(uuid.uuid4())
  response = Payment.cancel(
    payment_id,
    idempotence_key
  )
  if response.status == "canceled":
    print("Оплата отменена")
    return True
  else:
    print("Не удалось отменить оплату")
    return False