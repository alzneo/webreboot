window.onload = function() {
    document.getElementById("rebootBtn").onclick = reboot;
    document.getElementById("shutdownBtn").onclick = shutdown;
    document.getElementById("cancelBtn").onclick = cancel;
    document.getElementById("proceedBtn").onclick = proceed;
    document.getElementById("retryBtn").onclick = retry;
}

function switchToStep(step)
{
    for (let elem of document.getElementsByClassName("step")) {
        elem.classList.add("hidden");
    };
    document.getElementsByName(step)[0].classList.remove("hidden");
}

function reboot() {
    fetch("/power/reboot").then(response => {
        if (response.ok) {
            switchToStep("secondStep");
        } else {
            switchToStep("errorStep");
        }
    });
}

function shutdown() {
    fetch("/power/shutdown").then(response => {
        if (response.ok) {
            switchToStep("secondStep");
        } else {
            switchToStep("errorStep");
        }
    });
}

function cancel() {
    fetch("/power/cancel").then(response => {
        if (response.ok) {
            switchToStep("completeStep");
        } else {
            switchToStep("errorStep");
        }
    });
}

function proceed() {
    switchToStep("completeStep");
}

function retry() {
    switchToStep("firstStep");
}