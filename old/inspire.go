package main

type Motivation struct {
	quote  string
	source string
}
type Motivations []Motivation

// TODO: ini harus ditulis ulang sendiri
var motivations = []Motivation{
	{
		quote:  "Kebahagiaan hidupmu bergantung pada kualitas pikiranmu.",
		source: "Marcus Aurelius",
	},
	{
		quote:  "Kita lebih sering menderita dalam imajinasi daripada dalam kenyataan.",
		source: "Seneca",
	},
	{
		quote:  "Tidak masalah seberapa lambat kamu berjalan, asalkan kamu tidak berhenti.",
		source: "Confucius",
	},
	{
		quote:  "Satu-satunya cara untuk melakukan pekerjaan besar adalah dengan mencintai apa yang kamu kerjakan.",
		source: "Steve Jobs",
	},
	{
		quote:  "Imajinasi lebih penting daripada pengetahuan.",
		source: "Albert Einstein",
	},
	{
		quote:  "Segala sesuatu yang kamu inginkan ada di sisi lain ketakutan.",
		source: "George Addair",
	},
	{
		quote:  "Keberhasilan bukanlah akhir, kegagalan tidaklah fatal: keberanian untuk melanjutkanlah yang penting.",
		source: "Winston Churchill",
	},
	{
		quote:  "Kesulitan sering kali mempersiapkan orang biasa untuk nasib yang luar biasa.",
		source: "C.S. Lewis",
	},
	{
		quote:  "Percayalah kamu bisa, dan kamu sudah setengah jalan menuju ke sana.",
		source: "Theodore Roosevelt",
	},
	{
		quote:  "Waktumu terbatas, jadi jangan sia-siakan dengan menjalani hidup orang lain.",
		source: "Steve Jobs",
	},
	{
		quote:  "Cara terbaik untuk memprediksi masa depan adalah dengan menciptakannya.",
		source: "Alan Kay",
	},
	{
		quote:  "Lakukan apa yang kamu bisa, dengan apa yang kamu miliki, di mana pun kamu berada.",
		source: "Theodore Roosevelt",
	},
	{
		quote:  "Jadilah dirimu sendiri; orang lain sudah ada yang punya.",
		source: "Oscar Wilde",
	},
	{
		quote:  "Sesuatu selalu terlihat mustahil sampai hal itu selesai dilakukan.",
		source: "Nelson Mandela",
	},
	{
		quote:  "Perjalanan seribu mil dimulai dengan satu langkah.",
		source: "Lao Tzu",
	},
	{
		quote:  "Apa yang kita capai di dalam diri akan mengubah realitas luar.",
		source: "Plutarch",
	},
	{
		quote:  "Dia yang memiliki alasan untuk hidup dapat menanggung hampir semua cara untuk menjalaninya.",
		source: "Viktor Frankl",
	},
	{
		quote:  "Satu-satunya kebijaksanaan sejati adalah mengetahui bahwa kamu tidak tahu apa-apa.",
		source: "Socrates",
	},
	{
		quote:  "Bertindaklah seolah-olah apa yang kamu lakukan membuat perbedaan. Itu memang berpengaruh.",
		source: "William James",
	},
	{
		quote:  "Hidup adalah apa yang terjadi ketika kamu sibuk membuat rencana lain.",
		source: "John Lennon",
	},
	{
		quote:  "Kamu harus menjadi perubahan yang ingin kamu lihat di dunia.",
		source: "Mahatma Gandhi",
	},
	{
		quote:  "Jangan melihat jam; lakukan apa yang dilakukannya. Teruslah berjalan.",
		source: "Sam Levenson",
	},
	{
		quote:  "Pikiran adalah segalanya. Apa yang kamu pikirkan, itulah jadinya kamu.",
		source: "Buddha",
	},
	{
		quote:  "Hidup yang tidak diuji tidak layak untuk dijalani.",
		source: "Socrates",
	},
	{
		quote:  "Apakah kamu pikir kamu bisa atau kamu pikir kamu tidak bisa, kamu benar.",
		source: "Henry Ford",
	},
	{
		quote:  "Keunggulan bukanlah sebuah tindakan, melainkan sebuah kebiasaan.",
		source: "Aristotle",
	},
	{
		quote:  "Pikiran besar mendiskusikan ide; pikiran rata-rata mendiskusikan peristiwa; pikiran kecil mendiskusikan orang.",
		source: "Eleanor Roosevelt",
	},
	{
		quote:  "Satu-satunya batasan bagi realisasi kita akan hari esok adalah keraguan kita hari ini.",
		source: "Franklin D. Roosevelt",
	},
	{
		quote:  "Kreativitas adalah kecerdasan yang sedang bersenang-senang.",
		source: "Albert Einstein",
	},
	{
		quote:  "Dia yang berani adalah dia yang merdeka.",
		source: "Seneca",
	},
	{
		quote:  "Ubahlah lukamu menjadi kebijaksanaan.",
		source: "Oprah Winfrey",
	},
	{
		quote:  "Kapal di pelabuhan memang aman, tapi bukan untuk itu kapal dibuat.",
		source: "Grace Hopper",
	},
	{
		quote:  "Terkadang, orang-orang yang tidak pernah dibayangkan bisa melakukan hal-hal yang tidak terbayangkan.",
		source: "Alan Turing",
	},
	{
		quote:  "Belajarlah untuk berjualan. Belajarlah untuk membangun. Jika kamu bisa melakukan keduanya, kamu tidak akan terhentikan.",
		source: "Naval Ravikant",
	},
	{
		quote:  "Mengetahui orang lain adalah kecerdasan; mengetahui diri sendiri adalah kebijaksanaan sejati.",
		source: "Lao Tzu",
	},
	{
		quote:  "Dunia selalu terlihat lebih cerah saat kamu membuat sesuatu yang sebelumnya tidak ada.",
		source: "Neil Gaiman",
	},
}

func getDailyMotivation() Motivation {
	return motivations[getRandomDate()]
}

func getRandomMotivation() Motivation {
	return motivations[getRandomInt(len(motivations))]
}
