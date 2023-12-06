/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

const toast = new Toast(Toast.bottomRight);

jQuery(document).ready(function () {
    
    $("form").on("change, input", "input.danger-ic, select.danger-ic", function(){
        $(this).removeClass("danger-ic")
        $(this).closest(".input-group").removeClass("border-danger")
    })
    
    $("#validarDatos").click(function(){
        let formData = new FormData($("#ticketForm")[0])
        let formDataObject = {};
        formData.forEach((value, key) => (formDataObject[key] = value));
        formDataObject.negocioId = $("#negocioId").val();
        IFacturarDWR.validarRFC(formDataObject, (res)=>{
            if(res.result == "success"){
                changeStep(this)
                if(res.user){
                    fillSelect("regimenFiscal", res.fiscalRegimens )
                    fillSelect("usoCfdi", res.cfdiUses )
                    res.states && fillSelect("estado", res.states )
                    res.municipalities && fillSelect("municipio", res.municipalities )
                    $("select").val(0)
                    Object.entries(res.user).forEach(([key,value])=> {
                        let input = $(`[name=${key}]`)
                        value = value || (input.prop("tagName") == "INPUT" ? "" : 0)
                        input.val(value)
                    })
                    $(`[name=ticketId]`).val(res.ticketId)
                }else{
                    $("[name=rfc]").val($("[name=rfcTicket]").val())
                }
            }else{
                toast.addToast({type: Toast.error, body: res.message})
                Object.entries(res.campos).filter(([key,value]) => value == true).forEach(([key, _]) => {
                    $(`[name=${key}]`).addClass("danger-ic")
                    $(`[name=${key}]`).closest(".input-group").addClass("border-danger")
                })
            }
        })
    })
    
//    $("#userForm [name=rfc]").change(function(){
//        IFacturarDWR.rfcListas( $(this).val(), (res)=>{
//            console.log(res)
//        })
//    })
    
    $("#enviarFactura").click(function(){
        let ticketId = $(`[name=ticketId]`).val()
        let negocioId = $(`#negocioId`).val()
        let formData = new FormData($("#userForm")[0])
        let formDataObject = new ReceptorDWR();
        Object.seal(formDataObject)
        formData.forEach((value, key) => (formDataObject[key] = value));

        IFacturarDWR.obtenerFactura(ticketId, formDataObject, negocioId, (res) =>{
            if(res.result == "success"){
                changeStep(this)
                console.log(formDataObject)
                $("[name=correo]").val($("[name=correoEnviar]").val())
                $("#botonDescargaPdf").attr({href: res.pdf + ".pdf"})
                $("#botonDescargaXml").attr({href: res.pdf + ".xml"})
                $("#cfdiVistaVisor").attr("src", `${res.pdf}#toolbar=0&navpanes=0&scrollbar=0`)
            }else{
                toast.addToast({type: Toast.error, body: res.message})
                Object.entries(res.campos).filter(([key,value]) => value == true).forEach(([key, _]) => {
                    $(`[name=${key}]`).addClass("danger-ic")
                    $(`[name=${key}]`).closest(".input-group").addClass("border-danger")
                })
            }
        })
    })
    
    $("#inicioFacturador, #volverFacturador").click(function(){
        $(".facturacionForm").each((_,form)=>form.reset())
        changeStep(this)
    })
    
    $(".refillable").change(function(){
        switch($(this).attr("name")){
            case "codigoPostal":
                IFacturarDWR.searchList({postalCode: $(this).val()}, (res)=>{
                    fillSelect("estado", res.states, res.postalCode.StateCode)
                    fillSelect("municipio", res.municipalities, res.postalCode.MunicipalityCode)
                    $(`[name=pais`).val("MEX")
                    console.log(res.postalCode)
                })
                break
            case "pais":
                IFacturarDWR.searchList({states: $(this).val()}, (res)=>{
                    fillSelect("estado", res.states)
                    fillSelect("municipio")
                })
                break
            case "estado":
                IFacturarDWR.searchList({municipalities: $(this).val()}, (res)=>{
                    fillSelect("municipio", res.municipalities)
                })
                break
            case "rfc":
                
                IFacturarDWR.searchList({rfc: $(this).val()}, (res)=>{
                    fillSelect("usoCfdi", res.cfdiUses)
                    fillSelect("regimenFiscal", res.fiscalRegimens)
                })
                break;
        }
    })
    
    $(".uppercase").on("change, input", function(){
        let val = $(this).val()
        $(this).val(val.toUpperCase())
    })
});

function changeStep(actual = ""){
        let next = $(actual).attr("data-next")
        $(".facturacion-pasos").addClass("d-none")
        $(`#${next}`).removeClass("d-none")
        $("form .danger-ic, .input-group.border-danger").removeClass("danger-ic border-danger")
}

function fillSelect(select = "", optionList = [], val = null){
    let body = $(`[name=${select}]`)
    body.find("options").remove();
    let options = optionList.reduce((acc, curr) => { return `${acc}<option value=${curr.value || curr.Value}> ${curr.name || curr.Name} </option> `}, "<option value=0> Selecciona </option>")
    body.html(options);
    if(val){
        body.val(val);
    }
}
