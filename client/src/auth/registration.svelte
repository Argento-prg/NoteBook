<script>
    import { registrateService } from './authService'
    import { choiceUserStore } from './authStore'
    import { maxLengthName, maxLengthLogin, maxLengthPassword } from './authStore'
    let name = ''
    let login = ''
    let password = ''
    function register() {
        registrateService(name, login, password)
        .then((value) => {
            if (value) {
                choiceUserStore.update(() => false)
            }
        })
        .catch()
        .finally()
    }

</script>

<div class="form-register">
    <input 
        required
        placeholder="имя пользователя" 
        type="text" 
        maxlength={maxLengthName}
        bind:value={name}
    />
    <input 
        required
        placeholder="логин" 
        type="text" 
        maxlength={maxLengthLogin} 
        bind:value={login}
    />
    <input 
        required
        placeholder="пароль" 
        type="password" 
        maxlength={maxLengthPassword}
        bind:value={password}
    />
    <button on:click={register}>Зарегистрироваться</button>
</div>

<style>
    .form-register {
        background-color: antiquewhite;
        text-align: center;
        margin-right: auto;
        margin-left: auto;
        padding-top: 5%;
        border-bottom-left-radius: 20px;
        border-bottom-right-radius: 20px;
    }
    button {
        border-radius: 3%;
        border: none;
        background-color: #FF5C57;
        color: #FFFFFF;
    }
    button:hover {
        background-color: #FF2C3E;
        cursor: pointer;
    }
</style>