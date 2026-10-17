package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const pots = 10000               // сколько кастрюль ставим
const boilTime = 5 * time.Second // сколько закипает одна кастрюля

// boil варит одну кастрюлю. Плита греет сама, повар в это время свободен.
// В программе так выглядит любое ожидание: ответа сети, диска, базы данных.
func boil() {
	time.Sleep(boilTime)
}

// cookAll ставит n кастрюль сразу, каждую в своей горутине, и ждёт, пока сварятся все.
func cookAll(n int) {
	// TODO 1: заведи канал done и запусти n горутин.
	// Каждая горутина варит одну кастрюлю — boil(), — а потом кладёт в done пустой сигнал struct{}{}: «готово!»

	look() // все кастрюли на плите: смотрим, что творится на кухне

	// TODO 2: дождись всех кастрюль — забери из done ровно n сигналов.
}

// look печатает четыре числа: что сейчас на кухне. Код готовый, разбирать его не нужно.
func look() {
	time.Sleep(100 * time.Millisecond) // короткая пауза: пусть Go успеет раздать задачи поварам
	fmt.Println("ядра процессора  (рабочие места):", runtime.NumCPU())
	fmt.Println("процесс, его PID (заказ):        ", os.Getpid())
	fmt.Println("потоки ОС        (повара):       ", pprof.Lookup("threadcreate").Count())
	fmt.Println("горутины         (задачи):       ", runtime.NumGoroutine())
}

func main() {
	start := time.Now()
	cookAll(pots)
	fmt.Println(pots, "кастрюль сварились за", time.Since(start).Round(time.Second))
	fmt.Println("по очереди было бы", pots*boilTime)
}
