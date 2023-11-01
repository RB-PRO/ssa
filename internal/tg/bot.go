package tg

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RB-PRO/ssa/pkg/ssa"
	"github.com/RB-PRO/ssa/pkg/ssa2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Структура бота
type TG struct {
	Bot *tgbotapi.BotAPI
}

// Создаём бота
func NewBot(token string) (*TG, error) {
	bot, ErrNewBotAPI := tgbotapi.NewBotAPI(token)
	if ErrNewBotAPI != nil {
		return nil, ErrNewBotAPI
	}
	bot.Debug = false
	log.Printf("Авторизовался: %s", bot.Self.UserName)
	return &TG{bot}, nil
}
func BotStart2() {
	tg, ErrNewBot := NewBot("6720852434:AAFYUm2yplhsXWVKkXHa--o9tmkXZrNcpew")
	if ErrNewBot != nil {
		log.Panic(ErrNewBot)
	}
	tg.RangeUpdates()
}

// Слушаем сообшения в телеграме
func (tg *TG) RangeUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tg.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.VideoNote != nil {
			// Обработка видео-заметки
			videoNoteFile := update.Message.VideoNote.FileID
			videoNote, err := tg.Bot.GetFile(tgbotapi.FileConfig{FileID: videoNoteFile})
			if err != nil {
				log.Println("Ошибка при загрузке видео-заметки:", err)
				continue
			}

			// Сохраняем видео-заметку в локальный файл
			fileURL := "https://api.telegram.org/file/bot" + tg.Bot.Token + "/" + videoNote.FilePath
			FileDirection := update.Message.From.UserName + "_" + time.Now().Format("2006-01-02_15-04")
			FilePath := "TelegramVideoNote/" + FileDirection
			MakeDir(FilePath)
			err = saveVideoNoteLocally(fileURL, FilePath+"/"+FileDirection+".mp4")
			if err != nil {
				log.Println("Ошибка при сохранении видео-заметки:", err)
			}

			log.Printf("Пришло новое видео от пользователя %s. Видео в папке: %s",
				update.Message.From.UserName, FilePath)
			tg.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Видео получил. Достаю массив RGB(Не более 5 минут)"))

			// Массив RGB
			RGBs_float64, _ := ExtractRGB(FilePath+"/", FileDirection+".mp4")
			tg.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Достал RGB. Сейчас считаю pw"))
			tg.Bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(FilePath+"/"+"RGB.txt")))
			log.Printf("Рассчитано RGB")

			// Пульсовая волна
			pw, _ := CalcPW(RGBs_float64, FilePath+"/")
			tg.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Посчитал pw(Метод Cr)"))
			tg.Bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(FilePath+"/"+"pw.txt")))
			log.Printf("Рассчитано pw")

			////////////////////////////////////////////////////
			// SSA_tgbot(FilePath+"/", pw)
			// file := tgbotapi.FilePath(FilePath + "/smo.png")
			// tg.Bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, file))

			ssaAnalis, ErrNewSSA := ssa2.NewSSA(pw, ssa2.Setup{
				Cad:   23,
				Win:   1024,
				NPart: 20,
				FMi:   40.0 / 60.0,
				FMa:   240.0 / 60.0,
			})
			if ErrNewSSA != nil {
				tg.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrNewSSA.Error()))
				continue
			}
			ssa2.CreateLineChart(ssaAnalis.Pto_fMAX, FilePath+"/pto.png")
			file := tgbotapi.FilePath(FilePath + "/pto.png")
			tg.Bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, file))

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

		if _, err := tg.Bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// Обработка SSA-алгоритма
func SSA_tgbot(FilePath string, pw []float64) {
	s := ssa.New(FilePath + "Files/")
	s.Graph = false // Создавать графики
	s.Xlsx = true   // Сохранять в Xlsx
	s.Init(pw, []float64{})
	s.Spw_Form(pw) // Создать spw

	// # 1, 2, 3, 4
	s.SET_Form() // SSA - анализ сегментов pw

	// # 5
	// Оценка АКФ сингулярных троек для сегментов pw
	// Визуализация АКФ сингулярных троек для сегментов pw
	s.AKF_Form() // Оценка АКФ сингулярных троек для сегментов pw

	// # 6, 7
	// Огибающие АКФ сингулярных троек sET12 сегментов pw
	// Нормированные АКФ сингулярных троек sET12 сегментов pw
	s.Envelope()

	// # 8
	// Мгновенная частота нормированной АКФ сингулярных троек sET12 для сегментов pw
	s.MomentFrequency()

	// Дальнейшая обработка
	createLineChart(s.Tim, s.Smo_insFrc_AcfNrm, FilePath+"smo.png")
	// pw, _ := oss.Make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx
	// createLineChart()
}

///////////////////////////////////
///////////////////////////////////
///////////////////////////////////

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
