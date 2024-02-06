var languages = {
    "ru": {
        rebootBtn: "Перезагрузить",
        shutdownBtn: "Выключить",
        cancelMsg: "Есть около минуты, чтобы передумать",
        cancelBtn: "Отменить",
        proceedBtn: "Продолжить",
        completeMsg: "Готово",
        errorMsg: "Неизвестная ошибка!",
        retryBtn: "Повторить"
    }
};

document.querySelectorAll('[id]').forEach(elem => {
    var id = elem.id;
    var lang = navigator.language.substring(0, 2);
    elem.innerHTML = languages[lang][id];
});