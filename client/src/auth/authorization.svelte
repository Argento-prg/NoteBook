<script>
    import { authorizeService } from './authService'
    import { maxLengthLogin, maxLengthPassword } from './authStore'
    let login = ''
    let password = ''
    function enter() {
        authorizeService(login, password)
        .then()
        .catch()
        .finally()
    }
</script>

<div class="form-auth">
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
    <button on:click={enter}>Войти</button>
</div>

<style>
    .form-auth {
        background-color: khaki;
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