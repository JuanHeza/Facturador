/* global utilPortal, Intl, this */

//console.clear()
var facturaJS = null; 

class EvilToast {
    static success = "success";
    static info = "info";
    static data = "data";
    static error = "error";
    static warning = "warning";

    static topLeft = "NW";
    static topRight = "NE";
    static bottomLeft = "SW";
    static bottomRight = "SE";

    constructor(position=EvilToast.bottomRight) {
        this.templateSection = `<section id="evilToastSection" class="${position}"></section>`
        $("#evilToastSection").remove()
        $("body").append(this.templateSection)
        this.section = $("#evilToastSection");
        this.self = this
    }

    addToast({type=EvilToast.data, body="Hola", timer=5000, clear=false, enlace=null}={}) {
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

class EvilTable {
    static EvilData = {
        name: "dummy",
        title: "dummy",
        width: 0,
        visible: false,
        resize: false
    };
    
    static Formats = {
        date: "date",
        currency: "currency",
        icon: "icon"
    }
    
    static Aligns = {
        center: "text-center",
        left: "text-left",
        right: "text-right"
    }
    
    static filters = {
        text: "text",
        date: "date"
    }
    
    defaults = {
        visible: false,
        resize: false,
        color: "",
        function: "",
        icon: "",
        align: EvilTable.Aligns.left,
        search: null
    }
    
    url = ""
    dataType = "json"
    callType = "GET"
    cols = [EvilTable.EvilData]
    pager = '#perfilGridPager'
    height = 100
    rowNum = 10
    rowList = [10, 25, 50, 100]
    sortOrder = 'desc'
    sortName = 'desc'
    viewrecords = true
    body = "#evilTable"
    data = []
    pager = null
    pagination = true
    
    constructor(params={}) {
        Object.assign(this, params);
        let head = this.cols.reduce((acc,curr)=>{ return `${acc} 
            <th scope="col" class="text-center text-bold px-2 ${!curr.title ? "border-left-0 border-right-0" : ""}" style="${ curr.width ? `width: ${curr.width}px`  : `min-width: 80px` }"> ${curr.title} </th>`}
        ,"")
        let filters = this.cols.reduce((acc,curr)=>{ return {body:`${acc.body}
            <td class="${!curr.title ? "border-left-0 border-right-0" : ""} position-relative px-0 pb-4"> 
                    ${curr.search ? `<div class="d-inline-block position-absolute w-100" style="top: 1px;">
                            <input type="${curr.search}" class="form-control form-control-sm px-1 filterItem" name="${curr.name}"  placeholder="Buscar ${curr.title}">
                    </div>` : ""} 
            </td>`, count: curr.search ? acc.count+1 : acc.count}},{body:"", count:0})
        let template = $(this.bodyTemplate()).attr({id: this.body.replace("#","")})
        $(`${this.body}`).replaceWith(template)
        let local = this
        $(`${this.body} .paginationButton`).click(function(){
            let page = $(`${local.body} .pageDisplay`).val()
            let addition = $(this).attr("data-page")
            let val  = parseInt(page) + parseInt(addition)
            val = val < 1 ? 1 : val
            if(val != page){
                $(`${local.body} .pageDisplay`).val( val ).change()
            }
        })
        this.pager = `${this.body} .pageDisplay`
        $("#tableList").html(` <thead class="thead-light"> <tr> ${head} </tr> ${filters.count != 0 ? `<tr style="background: #f8f9fa"> ${filters.body} </tr> ` : ""}</thead><tbody></tbody>`)
    }
    
    bodyTemplate(){ 
        return `<section id="evilTable" class="overflow-auto rounded border">
            <article class="w-100 overflow-auto">
                <table id="tableList"  class="rounded table table-bordered table-hover table-sm " style="min-height: ${this.height}px; table-layout: fixed;"></table>
            </article>
            ${this.pagination ?  `<nav aria-label="Page navigation example">
                <ul class="pagination  justify-content-center">
                    <li class="cursor-pointer noselect page-item paginationButton" data-page="-1">
                        <i class="page-link material-icons border" aria-label="Previous">navigate_before</i>
                    </li>
                    <li class="col-1 p-0 page-item">
                        <input class="border-bottom border-top page-link pageDisplay text-black-50 text-bold text-center w-100" readonly="" disabled="" value="1">
                    </li>
                    <li class="cursor-pointer noselect page-item paginationButton" data-page="1">
                        <i class="page-link material-icons border" aria-label="Previous">navigate_next</i>
                    </li>
                </ul>
            </nav>` : "" }
        </section>`;
    }
    
    reloadTable(tableData = []) {
        console.log(this)
        this.data = tableData
        $(`${this.body} tbody tr`).remove()
        let filled = tableData.reduce((accRow, data)=>{
            let row = this.cols.filter(x => x.name).reduce((acc, curr)=>{
                curr = Object.assign({...this.defaults}, curr)
                return `${acc} <td class="px-2 ${curr.align} align-middle noselect"  style="overflow: hidden; white-space: nowrap; text-overflow: ellipsis;" title="${(curr.format ? this.format(curr.format, data[curr.name]) : data[curr.name]) || ""}"> ${(curr.format ? this.format(curr.format, data[curr.name]) : data[curr.name]) || ""} </td>`
            },"")
            row = this.cols.filter(x => !x.name).reduce((acc, curr)=>{
                let embedData = curr.data.reduce((accD,dat) => { return accD +`data-${dat}='${data[dat]}' `},"")
                return `${acc} <td class="material-icons d-table-cell border-left-0 border-right-0 text-center align-middle noselect cursor-pointer ${curr.color} ${curr.function}" ${embedData}> ${curr.icon} </td>`
            }, row)
            return `${accRow} <tr style="height: 45px"> ${row} </tr>`
        },"")
        $(`${this.body} tbody`).append(filled)
    }
    
    format(type = null, value = null){
        if(!type && !value){
            return value
        }
        switch(type){
            case EvilTable.Formats.date:
                return EvilTable.formatDate(value)
            case EvilTable.Formats.currency:
                return EvilTable.formatCurrency(value)
            case EvilTable.Formats.icon:
                return EvilTable.setIcon(value)
            default:
                return value;
        }
    }
    
    static formatCurrency(currency = 0){
        let format = Intl.NumberFormat('en-US', {
            style: "currency",
            currency: "USD"
        });
        return format.format(currency)
    }
    
    static formatDate(stringDate = "") {
        let date = new Date(stringDate)
        let day = String(date.getDate()).padStart(2, '0');
        let month = String(date.toLocaleString('es-MX', { month: 'short' })); // Months are 0-based
        let year = date.getFullYear();
        let hours = String(date.getHours()).padStart(2, '0');
        let minutes = String(date.getMinutes()).padStart(2, '0');

        return `${day}/${month}/${year} ${hours}:${minutes}`
    }
    
    static setIcon(status = ""){
        switch(status){
            case "active":
                return "<i class='align-middle m-0 material-icons noselect p-0 text-success' title='active'> check_circle </i>"
            case "canceled":
                return "<i class='align-middle m-0 material-icons noselect p-0 text-danger' title='canceled'> cancel </i>"
            default:
                return "<i class='align-middle m-0 material-icons noselect p-0 text-muted' title='unknow'> help </i>"
        }
    }
    
}

class Factura {
    constructor(params = {}){
        Object.assign(this, params);
        this.search =  new CfdiSearchDWR()
        this.search.cfdiType = 0
        let local = this
        $(`${this.table.body}`).on("change", ".filterItem", function(){
            let form = new CfdiSearchDWR()
            form.cfdiType = 0; 
            $(this).val($(this).val().trim())
            $(this).closest("tr").find(".filterItem").each((index, filter) => { form[$(filter).attr("name")] = $(filter).val().trim()})
             local.search =  form
             local.buscarCFDI()
        })
        
        $(this.table.pager).change(function(){
            console.log($(this).val())
            local.search.cfdiType = $(this).val()
            local.buscarCFDI()
        })
    }
    
    verCFDI(id = "", callback = this.fillVista){
        let search = new CfdiSearchDWR()
        search.id = id
        IFacturamaDWR.detailCfdi(search, callback)
    }
    
    crearCFDI(cfdi = null, cfdiItemList = null){
        IFacturamaDWR.facturar(cfdi, cfdiItemList, (res)=>{
            
            this.toast.addToast(res.id ==1 ? {type: EvilToast.success, body: res.valor || res.nombre } : { type: EvilToast.error, body: res.valor || res.nombre } )
            if(res.id ==1){
                $("#resetCFDI").click()
            }
            this.buscarCFDI()
        })
    }
    
    descargarCFDI(id = ""){
        IFacturamaDWR.downloadCfdi(id, (res)=>{
            this.toast.addToast(res.id == 1? {type: EvilToast.info, body: res.valor} : { type: EvilToast.error, body: "Error Inesperado" } )
            $("body").append(`<a id="dummyLink" href="${res.nombre.replace("c:","")}" download></a>`)
            $("#dummyLink")[0].click()
            $("#dummyLink").remove()
        }) 
    }

    cancelarCFDI(){
        let id = $("#deleteIdCFDI").val()       
        let motive = $("#deleteInputCFDI").val()
        let uuid = $("#replaceCfdi").val()
        if((motive == "04" && uuid != "" ) || motive != "04"){
            IFacturamaDWR.cancelCfdi(id, motive, uuid, (res)=>{
                this.toast.addToast(res ? {type: EvilToast.success, body: "CFDI cancelado" } : { type: EvilToast.error, body: "Error Inesperado" } )
            })
        }else{
            this.toast.addToast({ type: EvilToast.error, body: "Ingrese el CFDI de reemplazo" } )
        }
    }
    
    buscarCFDI(){
//        IFacturamaDWR.searchCfdi(this.search, (res)=>{this.table.reloadTable(res)})
    }
    
    fillVista(cfdi = null){
        if(cfdi){
            $("#cfdiVistaNombre").html(cfdi.taxName)
            $("#cfdiVistaFolio").html(cfdi.folio)
            $("#cfdiVistaRfcReceptor").html(cfdi.rfc)
            $("#cfdiVistaEmisor").html(cfdi.rfcIssuer)
            $("#cfdiVistaFecha").html(EvilTable.formatDate(cfdi.date))      
            $("#cfdiVistaEstatus").html(cfdi.status)
            let visor = $("#cfdiVistaVisor").clone().attr({src: `${cfdi.cfdiType}#toolbar=0&view=fitH&scrollbar=0` })
            $("#cfdiVistaVisor").replaceWith(visor);
            $(".cfdiCancelOptions").addClass("invisible")
            $("#downloadCFDI").removeClass("d-none")
            $("#deleteCFDI").addClass("d-none")
        }
        console.log(cfdi)
    }
    
    fillCancel(cfdi = null){
        if(cfdi){
            facturaJS.fillVista(cfdi)
            $("#downloadCFDI").addClass("d-none")
            $("#deleteCFDI").removeClass("d-none")
            $(".cfdiCancelOptions").removeClass("invisible")
            $("#deleteIdCFDI").val(cfdi.id)            
        }
    }    
    
    buscarCSD(){
//        IFacturamaDWR.SearchCsd("", (res)=>{this.tableCSD.reloadTable(res)})
    }
    
    createCSD(){
        let form = $('#formCSD')[0];
        let data = new FormData(form);
        let local = this
        $.ajax({
            type: "POST",
            enctype: 'multipart/form-data',
            url: "/cfdi/api/upload",
            data: data,
            processData: false,
            contentType: false,
            cache: false,
            timeout: 60000,
            success: function(data) {
                $("#result").text(data);
                console.log("SUCCESS : ", data);

                local.toast.addToast({
                    type: data.id == 1 ? EvilToast.success : EvilToast.error,
                    body: data.descripcion || data.nombre
                })
                local.buscarCSD()
                $("#resetCSD").click()
                $("#btnSubmit").prop("disabled", false);
            },

            error: function(e) {
                $("#result").text(e.responseText);
                console.log("ERROR : ", e);
                local.toast.addToast({
                    type: EvilToast.error,
                    body: e.responseText
                })
                $("#btnSubmit").prop("disabled", false);
            }
        });
    }
    
    removeCSD(rfc = null, callback = this.buscarCSD.bind(this)){
        if(rfc){
            let  data = new FormData();
//            data.append("rfc", rfc)
            let local = this
            IFacturamaDWR.removeCsd(rfc,(res)=>{
                local.toast.addToast({
                    type: res.id == 0 ? EvilToast.error : EvilToast.success,
                    body: res.valor || res.nombre
                }),
                callback()
            })
        }else{
            this.toast.addToast({
                type: EvilToast.error,
                body: "Seleccione un RFC valido"
            })
        }
    }
}

function addItemRow(){
    let template = $("#productRowTemplate").prop('content')
    let content = $(template).children().first().clone();
    $("#itemBox").append(content)
    $(".removeItemRow").toggleClass("invisible", $(".removeItemRow").length < 2)
}

function recalculateTotal(row = null){
    if (row){
        let cant = row.find(".cfdiItemCantidad").val() || 0
        let precio = row.find(".cfdiItemPrecio").val() || 0
        let impuesto = row.find(".cfdiItemPorcentajeImpuestos").val()
        let subtotal = cant * precio
        let total = subtotal + (subtotal * (impuesto/100))
        row.find(".cfdiItemSubtotal").val(subtotal.toFixed(2))
        row.find(".cfdiItemTotal").val(total.toFixed(2))
    }
}

function checkForm(form = null, target = null, isValid = null){
    let cfdiData = new FormData($(form)[0])
    if(!cfdiData && !target && !isValid){
        return false
    }
    let cfdiDummy = {}
    for (let [key, value] of cfdiData.entries()) {
        cfdiDummy[key] = value
        let input = $(`${form} [name=${key}]`)
        let havePattern = input.attr("pattern")
        let regex = new RegExp(havePattern)
        let isValid = havePattern ? regex.test(value) : true
        input.toggleClass("border-danger", !isValid)
        console.log(form, "-", key, "- Pattern Valid: ", havePattern ? regex.test(value) : "" )
    }
            
    Object.entries(cfdiDummy).filter(entry => !entry[1]).forEach(invalid => {
        let input = $(`${form} [name=${invalid[0]}]`)
        let isRequired = !!input.attr("required")
        input.toggleClass("border-danger", isRequired)
        isValid = isRequired ? false : isValid
        return true
    })
    Object.assign(target,cfdiDummy)
    return isValid
}

function checkItem(form = null, item = null, tax = null, isValid = null){
    let itemData = new FormData(form)
    let [taxesCfdi, itemCfdi] = [{},{}]

    for (var [key, value] of itemData.entries()) { 
        let input = $(form).find(`input[name='${key}']`)
        let isRequired = !!input.attr("required")
        input.toggleClass("border-danger", isRequired && !value)
        if(key == "name" || key == "rate" ){
            taxesCfdi[key] = value
        }else{
            itemCfdi[key] = value
        }
    }
    itemValid =  Object.entries(itemCfdi).every(entry => entry[1])
    taxValid = itemCfdi.taxObject != "02" || (itemCfdi.taxObject == "02" && taxesCfdi.rate != "")
    Object.assign(item, itemCfdi)
    Object.assign(tax, taxesCfdi)
    return isValid && taxValid && itemValid
}

class Test {
    personas = [{
           "rfc": "CACX7605101P8", "name": "XOCHILT CASAS CHAVEZ", "cp": "36257"
       },{
           "rfc": "FUNK671228PH6", "name": "KARLA FUENTE NOLASCO", "cp": "01160"
       },{
           "rfc": "IAÑL750210963", "name": "LUIS IAN ÑUZCO",   "cp": "825256"
       },{
           "rfc": "JUFA7608212V6", "name": "ADRIANA JUAREZ FERNANDEZ", "cp": "01160"
       },{
           "rfc": "KAHO641101B39", "name": "OSCAR KALA HAAK",  "cp": "76074"
       },{
           "rfc": "KICR630120NX3", "name": "RODRIGO KITIA CASTRO", "cp": "36246"
       },{
           "rfc": "MISC491214B86", "name": "CECILIA MIRANDA SANCHEZ",  "cp": "01010"
       },{
           "rfc": "RAQÑ7701212M3", "name": "ÑEVES RAMIREZ QUEZADA",    "cp": "78905"
       },{
           "rfc": "WATM640917J45", "name": "MARIA WATEMBER TORRES",    "cp": "43543"
       },{
           "rfc": "WERX631016S30", "name": "XAIME WEIR ROJO",  "cp": "01279"
       },{
           "rfc": "XAMA620210DQ5", "name": "ALBA XKARAJAM MENDEZ", "cp": "01219"
       },{
           "rfc": "XIQB891116QE4", "name": "BERENICE XIMO QUEZADA",    "cp": "40968"
       },{
           "rfc": "XOJI740919U48", "name": "INGRID XODAR JIMENEZ", "cp": "76028"
       },{
           "rfc": "EKU9003173C9",  "name": "ESCUELA KEMPER URGATE",    "cp": "42501"
       },{
           "rfc": "IIA040805DZ4",  "name": "INDISTRIA ILUMINADORA DE ALMACENES",   "cp": "62661"
       },{
           "rfc": "IVD920810GU2",  "name": "INNOVACION VALOR Y DESARROLLO",    "cp": "63901"
       },{
           "rfc": "IXS7607092R5",  "name": "INTERNACIONAL XIMBO Y SABORES",    "cp": "23004"
       },{
           "rfc": "JES900109Q90",  "name": "JIMENEZ ESTRADA SALAS",    "cp": "37161"
       },{
           "rfc": "KIJ0906199R1",  "name": "KERNEL INDUSTIA JUGUETERA",    "cp": "28971"
       },{
           "rfc": "OÑO120726RX3",  "name": "ORGANICOS ÑAVEZ OSORIO",   "cp": "40501"
       },{
           "rfc": "URE180429TM6",  "name": "UNIVERSIDAD ROBOTICA ESPAÑOLA",    "cp": "86991"
       },{
           "rfc": "XIA190128J61",  "name": "XENON INDUSTRIAL ARTICLES",    "cp": "76343"
       },{
           "rfc": "ZUÑ920208KL4",  "name": "ZAPATERIA URTADO ÑERI",    "cp": "34541"
       }
   ]
  
   fill(){
       let options = this.personas.reduce((acc, curr, ind)=>{ return `${acc} <option value=${ind}>${curr.name}</option>`},``)
       let local = this
       $("#testEmisor").html(options).change(function(){
           let persona = local.personas[$(this).val()]
           console.log(persona)
           $("#RfcEmisor").val(persona.rfc)
           $("#NombreEmisor").val(persona.name)
        })
       $("#testReceptor").html(options).change(function(){
           let persona = local.personas[$(this).val()]
           console.log(persona)
           $("#RfcReceptor").val(persona.rfc).trigger("input")
           $("#NombreReceptor").val(persona.name)
           $("#CodigoPostalReceptor").val(persona.cp)
        })
   }
}

jQuery(document).ready(function () {
    utilPortal.fnScrollPortal();

    let evilToast = new EvilToast(EvilToast.bottomRight);
    let evilTable = new EvilTable({cols: [
        {
            name: "status",
            title: "Estatus",
            format: EvilTable.Formats.icon,
            align: EvilTable.Aligns.center,
//            search: EvilTable.filters.text,
            width: 65
        },{
            name: "folio",
            title: "Folio",
            align: EvilTable.Aligns.center,
//            search: EvilTable.filters.text,
            width: 125
        },{
            name: "orderNumber",
            title: "Orden",
            align: EvilTable.Aligns.center,
            search: EvilTable.filters.text,
            width: 85
        },{
            name: "serie",
            title: "Serie",
            align: EvilTable.Aligns.center,
            search: EvilTable.filters.text,
            width: 65
        },{
            name: "rfcIssuer",
            title: "RFC Emisor",
            align: EvilTable.Aligns.center,
            search: EvilTable.filters.text,
            width: 125
        },{
            name: "date",
            title: "Fecha",
            format: EvilTable.Formats.date,
            align: EvilTable.Aligns.center,
            search: EvilTable.filters.date,
            width: 140
        },{
            name: "rfc",
            title: "RFC Receptor",
            align: EvilTable.Aligns.center,
            search: EvilTable.filters.text,
            width: 125
        },{
            name: "taxName",
            title: "Receptor",
            resizeable: true,
            search: EvilTable.filters.text,
            width: 260
        },{
            name: "total",
            title: "Total",
            format: EvilTable.Formats.currency,
            align: EvilTable.Aligns.center,
            width: 95
        },{
            name: "",
            title: "",
            color: "text-success",
            icon: "visibility",
            data: ["id"],
            function: "VerCfdi",
            width: 35
        },{
            name: "",
            title: "",
            color: "text-info",
            icon: "description",
            data: ["id"],
            function: "DescargarCfdi",
            width: 35
        },{
            name: "",
            title: "",
            color: "text-danger",
            icon: "highlight_off",
            data: ["id"],
            function: "CancelarCfdi",
            width: 35
        }
    ]})
    let csdTable = new EvilTable({cols: [
        {
            name: "rfc",
            title: "RFC Receptor",
            align: EvilTable.Aligns.center
        },{
            name: "",
            title: "",
            color: "text-danger",
            icon: "highlight_off",
            function: "eliminarCsd",
            data: ["rfc"],
            width: 50
        }
    ], body: "#csiList", pagination: false})
    
    facturaJS = new Factura({table: evilTable, toast: evilToast, tableCSD: csdTable})
    
    facturaJS.buscarCFDI()
    facturaJS.buscarCSD()
    
    let test = new Test()
    test.fill()
    
    addItemRow()

    $("#acceptCSD").off("click").click(function() {
        facturaJS.createCSD();
    })
    
    $("#resetCSD").off("click").click(function() {
        $(".csdInput").val(null)
    })
    
    $("input[type=file].csdInput").off("change").change(function() {
        $(`#${$(this).attr("id")}label .csdPlaceholder`).val($(this).val().replace(/C:\\fakepath\\/i, ''))
    })
    
    $("#deleteCFDI").click(function(){
        facturaJS.cancelarCFDI();
    })
    
    $("#volverCFDI").click(function(){
        $("#deleteInputCFDI").prop('selectedIndex',0);
        $("#replaceCfdi").val("").attr({disabled: true})
    })
    
    $("#deleteInputCFDI").change(function(){
        let status = $(this).val() != "01"
        $("#replaceCfdi").attr({disabled: status, required: !status}).val("")
        $("#deleteCFDI").attr({disabled : !status}).toggleClass("disabled", !status)
    })
    
    $("#replaceCfdi").on("input", function(){
        let pattern = /^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i;
        let valid = pattern.test($(this).val().trim());
        $("#deleteCFDI").attr({disabled : !valid}).toggleClass("disabled", !valid)
    })
    
    $("#RfcReceptor").on("input", function(){
        let rfcMoral = /^[a-zA-Z,Ñ,ñ,&]{3}\d{6}[A-V1-9]{1}[A-Z\d]{2}$/i;
        let rfcFisica = /^[a-zA-Z,Ñ,ñ,&]{4}\d{6}[A-V1-9]{1}[A-Z\d]{2}$/i;
        let val = $(this).val()
        let cfdiUse = $("input[name='UsoCfdi']")
        switch(true){
            case rfcFisica.test(val):
                cfdiUse.val('natural').change()
                break;
            case rfcMoral.test(val):
                cfdiUse.val('moral').change()
                break;
            default:
                cfdiUse.val('null').change()
        }
    })
    
    $(facturaJS.table.body).on("click", ".VerCfdi", function(){
        facturaJS.verCFDI($(this).data("id"), facturaJS.fillView)
        console.log(`%c ${$(this).data("id")} `, "color:green")
    })
    
    $(facturaJS.table.body).on("click", ".DescargarCfdi", function(){
        facturaJS.descargarCFDI($(this).data("id"), facturaJS.descargarArchivo)
        console.log(`%c ${$(this).data("id")} `, "color:blue")
    })
    
    $("#downloadCFDI").click(function(){
        facturaJS.descargarCFDI($("#deleteIdCFDI").val(), facturaJS.descargarArchivo)
    })
    
    $(facturaJS.table.body).on("click", ".CancelarCfdi", function(){
        facturaJS.verCFDI($(this).data("id"), facturaJS.fillCancel)
        console.log(`%c ${$(this).data("id")} `, "color:red")
    })

    $(facturaJS.tableCSD.body).on("click", ".eliminarCsd", function(){
        facturaJS.removeCSD($(this).data("rfc"))
        console.log(`%c ${$(this).data("id")} `, "color:red")
    })
    
    $("input[name='UsoCfdi']").change(function(){ 
        let val = $(this).val() 
        $("#UsoCfdiSelectNatural").toggleClass("d-block", val == "natural")
        $("#UsoCfdiSelectMoral").toggleClass("d-block", val == "moral")
        $("#UsoCfdiDummy").toggleClass("d-block", val == "null")
    })
  
    $("#itemBox").on("change", ".cfdiImpuestosSelect", function() {
        let val = $(this).val()
        let row = $(this).closest(".row")
        row.find(".cfdiImpuestosInput").toggleClass("invisible", val != "02").removeClass("border-danger")
        row.find("input.cfdiImpuestosInput").val("").trigger("input").prop({required: val == "02"})
        row.find("select.cfdiImpuestosInput").prop("selectedIndex", 0);
    })
    
    $("#itemBox").on("input", ".cfdiItemCantidad, .cfdiItemPrecio", function() {
        recalculateTotal($(this).closest(".row"))
    })
    
    $("#itemBox").on("input", "input.cfdiImpuestosInput", function() {
        let val = $(this).val()
        if(val > 100){
            val = 100
        }
        if(val < 0){
            val = 0
        }
        $(this).val(val)
        recalculateTotal($(this).closest(".row"))
    })

    $("body").off("click", ".addItemRow").on("click", ".addItemRow", function(){
        addItemRow()
    })
    
    $("body").off("click", ".removeItemRow").on("click", ".removeItemRow", function(){
        $(this).closest(".row").remove()
        $(".removeItemRow").toggleClass("invisible", $(".removeItemRow").length < 2)
    })

    $("#acceptCFDI").off("click").on("click", function(){ 
        console.clear()
        let cfdiItemList = []
        let isValid = true
        
        let Cfdi = new CfdiDWR()
        isValid = checkForm("#formCFDI", Cfdi, isValid)
        
        let CfdiEmisor = new CfdiEmisorDWR()
        isValid = checkForm("#CfdiIssuerForm", CfdiEmisor, isValid)
        
        let CfdiReceptor = new CfdiReceptorDWR()
        isValid = checkForm("#CfdiReciverForm", CfdiReceptor, isValid)
        
        $(".rowItem").each((_,item)=>{
            let itemCfdi = new CfdiItemDWR()
            let taxesCfdi = new CfdiTaxDWR()
            isValid = checkItem(item, itemCfdi, taxesCfdi, isValid)
            if(!isValid){
                return false;
            }
            itemCfdi.unitCode = $("#cfdiGeneralUnits").val()
            itemCfdi.taxes = itemCfdi.taxObject == "02" ? [] : null
            if(itemCfdi.taxObject == "02"){
                taxesCfdi.rate /= 100
                taxesCfdi.base = parseFloat(itemCfdi.subtotal)
                taxesCfdi.total =  (taxesCfdi.base * taxesCfdi.rate).toFixed(2)
                itemCfdi.taxes.push(taxesCfdi)
            }
            
            cfdiItemList.push(itemCfdi)
        })
        
        let useCfdi = $("input[name='UsoCfdi']:checked").val()
        if(useCfdi != "null" && isValid){
            CfdiReceptor.cfdiUse = useCfdi == "moral" ? $("#UsoCfdiSelectMoral").val() : $("#UsoCfdiSelectNatural").val()
            
            Cfdi.issuer = CfdiEmisor
            Cfdi.receiver = CfdiReceptor
            Cfdi.currencyExchangeRate = Cfdi.currency == "MXN" ? null : 35.0
            console.log(Cfdi, cfdiItemList)
            facturaJS.crearCFDI(Cfdi, cfdiItemList)
        }else{
            $(".border-danger")[0].scrollIntoView({ block: "center", behavior: "smooth" })
            facturaJS.toast.addToast({type: EvilToast.error, body: "Revise los datos"})
        }
    })
    
    $("#resetCFDI").on("click", function(){
         $("#formCFDI")[0].reset()
         $("#CfdiIssuerForm")[0].reset()
         $("#CfdiReciverForm")[0].reset()
         $("[name=UsoCfdi]").val("null").change()
         $(".rowItem").remove()
         addItemRow()
    })
    
    $("#cfdiGeneralLogo").change(function(){
        $("#cfdiGeneralLogoPreview").attr("src", $(this).val())
    })
    
    $("form").on("change", ".border-danger", function(){
        $(this).removeClass("border-danger")
    })
    /*
    let timer = null
    $(".hoverArrow").mouseover(function() {
        let [p1, w1] = [$(this).position(), ($(this).outerWidth() / 2)] 

        let [x1, y1] = [p1.left + w1, p1.top]

        let [x2, y2] = [x1, y1 - 30]

        let [p2, w2] = [$("#info-box").position(), ($("#info-box").outerWidth())] 

        let [x3, y3] = [p2.left + w2, y2]

        let [x4, y4] = [x3, p2.top]

        let newLine = document.createElementNS('http://www.w3.org/2000/svg','polyline');
        newLine.setAttribute('id','line2');
        newLine.setAttribute('points', `${x1},${y1} ${x2},${y2} ${x3},${y3} ${x4},${y4}`);
        newLine.setAttribute('fill', "none");
        newLine.setAttribute('stroke', "black");

        timer = setTimeout(() => {
            $("#info-box > div").addClass("h-100")
        }, 800);
        $("#line2fading").remove();       
        $("#map").append(newLine);      
    }).mouseout(function() {
//            setTimeout(() => {
            $("#line2").attr("id", "line2fading")
//            }, 500);
        clearTimeout(timer);
        $("#info-box > div").removeClass("h-100")
    });
    */
});