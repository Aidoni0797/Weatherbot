from ttoken import Token
import telebot
from telebot import types

bot = telebot.TeleBot(Token)

name = ''
age = 0

@bot.message_handler(content_types=['text'])
def getTextMessage(message):
    if message.text == 'hello':
        bot.send_message(message.chat.id,'hello, what`s your name')
        bot.register_next_step_handler(message, getname)


def getname(message):
    global name
    name = message.text
    bot.send_message(message.chat.id,'Your name is ' + name)
    bot.send_message(message.chat.id, 'hello, how old are you')
    bot.register_next_step_handler(message, getage)

def getage(message):
    global age
    age = message.text
    bot.send_message(message.chat.id,'Your age is ' + age)
    keyboard = types.InlineKeyboardMarkup()
    keyYes = types.InlineKeyboardButton('Yes',callback_data='yes')
    keyNo = types.InlineKeyboardButton('No',callback_data='no')
    keyboard.add(keyYes,keyNo)
    bot.send_message(message.chat.id,'Your name is '+name
                     + ' and your age is ' + age + '?',reply_markup=keyboard)
@bot.callback_query_handlers(func=lambda call:True)
def callBackHandler(call):
    if call.data =='yes':
        bot.send_message(call.message.chat.id,'Your data is saved')
    elif call.data =='no':
        bot.send_message(call.message.chat.id,'Whats your name')
        bot.register_next_step_handler(call.message,getname)
bot.polling(none_stop=True,interval=0)