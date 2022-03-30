<script>

</script>

<div class="footer">
    <div class="aboutApp">
        О приложении
        <p>
            Приложение представляет собой расширяемую записную книжку.
            Архитектура клиент-сервер. Используемые технологии: Go, SvelteJS, PostgreSQL.
        </p>
    </div>
    <div class="aboutUs">
        О нас
        <p>
            Приложение было разработано студентом 4 курса ИНБс 
            Бровцыным Олегом в качестве практической работы по дисциплине 
            СУБД.
        </p>
    </div>
</div>

<style>
    .footer {
        background-color: rgb(126, 247, 206, 0.5);
        padding: 0;
        margin: 0;
        height: 10%;
        width: 100%;
        display: flex;
    }
    .aboutApp {
        width: 50%;
        display: block;
        text-align: center;
        border-right: inset black 1px;
    }
    .aboutUs {
        width: 50%;
        display: block;
        text-align: center;
        border-left: inset black 1px;
    }
    p {
        margin: 0;
        padding: 0;
        font-size: 12px;
    }
</style>
