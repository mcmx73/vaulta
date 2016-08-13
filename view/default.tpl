<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>VAULTA Test Console</title>

    <!-- Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"
          integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
    <link href="console.css" rel="stylesheet">
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

</head>
<body>

<nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container-fluid">
        <div class="navbar-header">
            <a class="navbar-brand" href="#" onclick="return(false);">Vaulta API Test console</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav navbar-right">
                <li><a href="/">Reload dashboard</a></li>
            </ul>


        </div>
    </div>
</nav>

<div class="container-fluid">
    <div class="row">
        <div class="col-sm-12 col-md-12 main">
            <div class="row" style="display: none">
                <div class="col-md-12" role="group">
                    <label class="form-signin-heading">HTTP Server Address</label>
                    <div class="input-group"><input type="text" class="form-control" id="http_server" name="http_server"
                                                    placeholder="HTTP Address" value="http://vaulta.local"
                                                    onblur="setCookie('http_server', this.value);">
			              <span class="input-group-btn">
				              <button class="btn  btn-primary" type="button" onclick="pingHTTP();return(false)">Ping
                                  HTTP Server
                              </button>
			              </span>
                    </div>
                </div>
            </div>
            <div class="row" id="sign_block">
                <hr/>
                <div class="col-md-6" role="group">
                    <div class=" panel panel-default">
                        <div class="panel-heading">Vaulta Save data</div>
                        <div class="panel-body">
                            <div class="form-group">
                                <label class="form-signin-heading">Data for save <span id="result" style="color: #428bca"></span></label>
                                <textarea type="text" id="data_block_save" name="data_block"
                                          class="form-control" rows="10"
                                          placeholder="The Ultimate Question of Life, the Universe, and Everything"></textarea>
                            </div>
                            <div class="form-group">
                                <label class="form-signin-heading">Data Link</label>
                                <input type="text" id="data_link" name="data_link"
                                       class="form-control"
                                       placeholder="..."
                                       value="" readonly>
                            </div>
                            <div class="form-group">
                                <label class="form-signin-heading">Data Key</label>
                                <input type="text" id="data_key" name="data_key"
                                       class="form-control"
                                       placeholder="..."
                                       value="" readonly>
                            </div>

                        </div>
                        <div class="panel-body col-md-12">
                            <div class="form-group">
                                <button class="btn  btn-primary btn-block" type="button"
                                        onclick="vaultaSave()">Save
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
                <form class="col-md-6" role="group">
                    <div class=" panel panel-default row">
                        <div class="panel-heading">Vaulta Read data</div>
                        <div class="panel-body col-md-12">
                            <div class="form-group">
                                <label class="form-signin-heading">Data Link</label>
                                <input type="text" id="data_link_r" name="data_key"
                                       class="form-control"
                                       placeholder="..."
                                       value="">
                            </div>
                            <div class="form-group">
                                <div class="form-group">
                                    <label class="form-signin-heading">Data Key</label>
                                    <input type="text" id="data_key_r" name="data_key"
                                           class="form-control"
                                           placeholder="..."
                                           value="">
                                </div>
                            </div>
                            <div class="form-group">
                                <label class="form-signin-heading">data_block</label>
                                <textarea id="data_block" name="data_block"
                                          class="form-control"
                                          placeholder="..."
                                          rows="10" readonly></textarea>
                            </div>
                        </div>
                        <div class="panel-body col-md-12">
                            <div class="form-group">
                                <button class="btn  btn-primary btn-block" type="button"
                                        onclick="vaultaLoad()">Load
                                </button>
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
</div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"
        integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS"
        crossorigin="anonymous"></script>
<script src="/vaulta.js"></script>
</body>
</html>
