function vaultaSave() {
    server = "http://vaulta.local";
    addr = server + "/api"
    try {
        $.ajax({
            type: "POST",
            url: addr,
            data: JSON.stringify({
                data_block: document.getElementById('data_block_save').value,
            }),
            dataType: "json",
            contentType: "application/json; charset=utf-8",
            complete: function (data) {
                console.info(data.responseJSON)
                document.getElementById("data_link").value = data.responseJSON.block_index
                document.getElementById("data_link_r").value = data.responseJSON.block_index
                document.getElementById("data_key").value = data.responseJSON.key_index
                document.getElementById("data_key_r").value = data.responseJSON.key_index
                document.getElementById("result").innerHTML = data.responseJSON.Result
            }
        });
    } catch (netError) {
        console.error(netError)
    }
}

function vaultaLoad() {
    server = "http://vaulta.local";
    addr = server + "/api/" + document.getElementById("data_link_r").value + "/" + document.getElementById("data_key_r").value
    try {
        $.ajax({
            type: "GET",
            url: addr,
            dataType: "json",
            contentType: "application/json; charset=utf-8",
            complete: function (data) {
                document.getElementById("data_block").value = data.responseJSON.data_block
                console.info(data.responseJSON)
            }
        });
    } catch (netError) {
        console.error(netError)
    }
}