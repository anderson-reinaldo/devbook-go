
$('#register-form').on('submit',registerUser)

function registerUser(event){
    event.preventDefault();

     const data = {
        nome: $('#nome').val(),
        nick: $('#nick').val(),
        email: $('#email').val(),
        senha: $('#senha').val(),
        confirmarSenha: $('#confirmar-senha').val(),

     }

     if(!data.nome || !data.nick || !data.email || !data.senha || !data.confirmarSenha){
        alert("Todos os campos são obrigatorios!")
        return
     }


    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("As senhas não coincidem!")
        return
    }

    delete data.confirmarSenha

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data
    }).done(res => {
        alert("usuario cadastrado com sucesso!")
    })
    .fail(err => {
        alert(err.responseJSON.error)
    })
}