package tg

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotStart() {
	bot, err := tgbotapi.NewBotAPI("6720852434:AAFYUm2yplhsXWVKkXHa--o9tmkXZrNcpew")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.VideoNote != nil {
			// Обработка видео-заметки
			videoNoteFile := update.Message.VideoNote.FileID
			videoNote, err := bot.GetFile(tgbotapi.FileConfig{FileID: videoNoteFile})
			if err != nil {
				log.Println("Ошибка при загрузке видео-заметки:", err)
				continue
			}

			// Сохраняем видео-заметку в локальный файл
			fileURL := "https://api.telegram.org/file/bot" + bot.Token + "/" + videoNote.FilePath
			FileDirection := update.Message.From.UserName + "_" + time.Now().Format("15-04_02-01-2006")
			FilePath := "TelegramVideoNote/" + FileDirection
			MakeDir(FilePath)
			err = saveVideoNoteLocally(fileURL, FilePath+"/"+FileDirection+".mp4")
			if err != nil {
				log.Println("Ошибка при сохранении видео-заметки:", err)
			}

			log.Printf("Пришло новое видео от пользователя %s. Видео в папке: %s",
				update.Message.From.UserName, FilePath)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Видео получил. Достаю массив RGB(Не более 5 минут)"))

			// Массив RGB
			RGBs_float64, _ := ExtractRGB(FilePath+"/", FileDirection+".mp4")
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Достал RGB. Сейчас считаю pw"))
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(FilePath+"/"+"RGB.txt")))
			log.Printf("Рассчитано RGB")

			// Пульсовая волна
			CalcPW(RGBs_float64, FilePath+"/")
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Посчитал pw(Метод Cr)"))
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(FilePath+"/"+"pw.txt")))
			log.Printf("Рассчитано pw")
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Привет. Я - бот, который умеет в дистанционную фотоплетизмографию. Если не понятно, что со мной делать - /help."
		case "help":
			msg.Text = "Запишите мне видео в кружочке, а я Вам покажу всё, что умею" +
				"Список команд:\n" +
				" /start - Общая информация\n" +
				" /help - Помощь\n" +
				" /gratitude - Благодарности\n" +
				" /developer - Разработчик бота\n"
		case "gratitude":
			msg.Text = "Благодарности:\n@rkhn_maria - Ведущий научный сотрудник МГТУ им. Баумана"
		case "developer":
			msg.Text = "Разработчик бота - @RB_PRO"
		default:
			msg.Text = "Я не знаю такую команду. Попробуй /help"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func saveVideoNoteLocally(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
