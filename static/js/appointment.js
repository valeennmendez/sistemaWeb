console.log("Conectado...")

function GetAllAppointments(){

    fetch(`https://sistemaweb-production.up.railway.app/appointments`,{
        method: "GET",
        credentials: "include"
    })
        .then(response =>{
            if(!response.ok){
                console.error("ERROR")
            }
            return response.json()
        })
        .then(data =>{
            const tableBody = document.querySelector(".tabla tbody")
            tableBody.innerHTML = "";

            data.forEach(appointment =>{
                const row = document.createElement("tr")
                
                const fecha = new Date(appointment.Fecha)
                const fechaFormateada = fecha.toLocaleDateString()

                row.innerHTML = `
                <td>${appointment.Paciente.FullName}</td>
                <td>${appointment.Paciente.Dni}</td>
                <td>${fechaFormateada}</td>
                <td>${appointment.Hora}</td>
                `
                tableBody.appendChild(row)
            })
        })
        .catch(error => console.error("error: ",error))
}

function CancelAppointments(){
    const button = document.getElementById("btnCancel")

    if(button){
        button.addEventListener("click", function(e){
            
            fetch(`https://sistemaweb-production.up.railway.app/cancel-appointment`)
        })
    }
}

function SearchPatientsForm(){
    const inputSearch = document.getElementById("searchPatient")

    inputSearch.addEventListener("input", function(){
        const query = this.value;

        if(query.length > 2){
            fetch(`https://sistemaweb-production.up.railway.app/search-patient?p=${encodeURIComponent(query)}`)
                .then(response => response.json())
                .then(data =>{
                    const resultDropdown = document.getElementById("resultsDropdown")
                    resultDropdown.innerHTML = ""

                    data.forEach(patient =>{
                        const option = document.createElement("option")
                        option.value = patient.FullName
                        option.text = patient.FullName
                        resultDropdown.appendChild(option)
                    })
                })
                .catch(error => console.error("error", error))
        }

    })
}

document.addEventListener("DOMContentLoaded", function(e){

    GetAllAppointments()

    SearchPatientsForm()

})
