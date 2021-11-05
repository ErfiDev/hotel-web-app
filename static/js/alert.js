function Notify(msg = "" , icon = "info"){
    const toast = Swal.mixin({
        toast: true,
        title: msg,
        icon: icon,
        position: "bottom-right",
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast)=> {
            toast.addEventListener("mouseenter" , Swal.stopTimer)
            toast.addEventListener("mouseleave" , Swal.resumeTimer)
        }
    })
    toast.fire({})
}



const form = document.querySelector(".make-reservation-form")

form.addEventListener("submit" , (e)=>{
    e.preventDefault()

    Notify("salam")
})

