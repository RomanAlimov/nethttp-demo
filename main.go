package main

import ( // необходимые пакеты
	"fmt"
	"log"
	"net/http"
)

func logging(next http.Handler) http.Handler { // создал функцию обработчика и вернул ее новой функцией обработчика
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //возврат новой функции в ней же запрос и завершение
		log.Println("Запрос:", r.Method, r.URL.Path) // откладка в результате которого выводится ссылка куда зашел пользователь
		next.ServeHTTP(w, r)                         // передача инфы
		log.Println("Завершение отправки запроса.")  // откладка в терминал
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Обработано") //напечатал
}

func main() {
	mux := http.NewServeMux()              //переменная
	mux.HandleFunc("/hello", helloHandler) // адрес в локалке

	wrappedMux := logging(mux) // передача переменной

	http.ListenAndServe(":8080", wrappedMux) // наш сервер на локалке
}
