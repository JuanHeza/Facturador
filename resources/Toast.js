/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */



class Toast {
    static success = "success";
    static info = "info";
    static data = "data";
    static error = "error";
    static warning = "warning";

    static topLeft = "NW";
    static topRight = "NE";
    static bottomLeft = "SW";
    static bottomRight = "SE";

    constructor(position=Toast.bottomRight) {
        this.templateSection = `<section id="evilToastSection" class="${position}"></section>`
        $("#evilToastSection").remove()
        $("body").append(this.templateSection)
        this.section = $("#evilToastSection");
        this.self = this
    }

    addToast({type=Toast.data, body="Hola", timer=5000, clear=false, enlace=null}={}) {
        let icons = {
            success: "check_circle",
            error: "cancel",
            data: "explore",
            info: "info",
            warning: "warning",
            link: "language"
        }
        let tag = enlace ? "a" : "div"
        let template = `<${tag} class="${type} evilToast show">
            <div class="body">
                <i class="material-icons">${icons[type]}</i>
                <p class="text">${body}</p>
            </div>
        </${tag}>`

        if (clear) {
            this.removeToast()
        }

        let toast = $(template)
        $("#evilToastSection").append(toast)

        let self = this.self
        let destroy = timer == 0 ? null : this.timerRemoveToast(toast, timer)

        $(toast).mouseover(function() {
            clearTimeout(destroy)
        }).mouseout(function() {
            destroy = self.timerRemoveToast(toast, timer)
        });

        $("body").off("click", ".evilToast").on("click", ".evilToast", function() {
            self.removeToast($(this))
        })
    }

    removeToast(toast=$(".evilToast")) {
        toast.addClass("hide")
        setTimeout(function() {
            toast.remove()
        }, 1000)
    }

    timerRemoveToast(toast=null, timer=5000) {
        if (toast) {
            return setTimeout(()=>{
                this.removeToast(toast)
            }, timer)
        }
    }
}