<script>
    import Authorization from './authorization.svelte'
    import Registration from './registration.svelte'
    import { choiceUserStore } from './authStore'
    let choiceUser
    choiceUserStore.subscribe(value => { choiceUser = value })
</script>

<div class="auth">
    <span id="logo">NoteBook</span>
    <div class="main-form-auth">
        <div class="choice">
            <button 
                on:click={() => { choiceUserStore.update(value => false) }}
                class="auth-btn"
            >
                Авторизация
            </button>
            <button 
                on:click={() => { choiceUserStore.update(value => true) }}
                class="reg-btn"
            >
                Регистрация
            </button>
        </div>
        {#if !choiceUser}
            <Authorization />
        {:else}
            <Registration />
        {/if}
    </div>
</div>

<style>
    .auth {
        width: 100%;
        height: 90%;
        background: linear-gradient(30deg, #69b7eb, #b3dbd3);
    }
    #logo {
        font-size: 40px;
        text-align: center;
        display: block;
        padding-top: 5%;
    }
    .main-form-auth {
        width: 250px;
        margin: 0 auto;
        padding: 2%;
    }
    .choice {
        display: flex;
        justify-content: center;
    }
    button:hover {
        cursor: pointer;
    }
    button {
        width: 50%;
        display: inline-block;
        border: none;
        margin-bottom: 0;
        border-radius: 0;
    }
    .auth-btn {
        background-color: khaki;
    }
    .auth-btn:hover {
        background-color: #FFEC60;
    }
    .reg-btn {
        background-color: antiquewhite;
    }
    .reg-btn:hover {
        background-color: rgb(245, 192, 122);
    }
</style>