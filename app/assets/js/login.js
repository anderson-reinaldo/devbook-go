
$('#login-form').on('submit',loginUser)

function loginUser(event){
    event.preventDefault();

     const data = {
        email: $('#email').val(),
        senha: $('#senha').val(),
     }

     console.log(data,'--data')
     if(!data.email || !data.senha){
        alert("Todos os campos são obrigatorios--!")
        return
     }

    $.ajax({
        url: "/login",
        method: "POST",
        data
    }).done(() => {
        window.location = "/home"
    })
    .fail(err => {
        alert(err.responseJSON.error)
    })
}