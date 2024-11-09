package main

import (
	"html/template"
	"log"
	"net/http"
)

type Part struct {
	Image string
	Title string
	Price string
}

type Machine struct {
	Image string
	Title string
	Price string
}

type MachineList struct {
	Machines []Machine
	Parts    []Part
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static", fileServer))

	log.Println("Starting Server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error lsitening")
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	parts := []Part{
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹3000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/wheel.png", Title: "Wheel", Price: "₹5000"},
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹6000"},
		{Image: "/web/static/images/bearing.png", Title: "Bearing", Price: "₹6000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/wheel.png", Title: "Wheel", Price: "₹5000"},
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹6000"},
		{Image: "/web/static/images/bearing.png", Title: "Bearing", Price: "₹6000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹3000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/wheel.png", Title: "Wheel", Price: "₹5000"},
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹6000"},
		{Image: "/web/static/images/bearing.png", Title: "Bearing", Price: "₹6000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		{Image: "/web/static/images/wheel.png", Title: "Wheel", Price: "₹5000"},
		{Image: "/web/static/images/bucket.png", Title: "Bucket", Price: "₹6000"},
		{Image: "/web/static/images/bearing.png", Title: "Bearing", Price: "₹6000"},
		{Image: "/web/static/images/excavator.png", Title: "Excavator", Price: "₹25000"},
		// Add more parts as needed
	}

	machine := []Machine{
		{Image: "/web/static/images/ex200.png", Title: "EX 200 - 2012 ", Price: "₹1,30,000/month"},
		{Image: "/web/static/images/ex200.png", Title: "EX 200 -2011 ", Price: "₹1,30,000/month"},
		{Image: "/web/static/images/machine3.png", Title: "EX 200 - 2016 ", Price: "₹1,40,000/month"},
	}

	machines := MachineList{Machines: machine, Parts: parts}

	t := template.Must(template.New("home").ParseFiles("./web/index.html", "./web/base.html", "./web/cards.html", "./web/machines.html"))

	t.ExecuteTemplate(w, "home", machines)
}
