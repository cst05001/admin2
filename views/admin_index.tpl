<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <!-- 新 Bootstrap 核心 CSS 文件 -->
        <link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.0/css/bootstrap.min.css">
        <!-- 可选的Bootstrap主题文件（一般不用引入） -->
        <link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.0/css/bootstrap-theme.min.css">
        <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
        <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
        <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
        <script src="http://cdn.bootcss.com/bootstrap/3.3.0/js/bootstrap.min.js"></script>
    </head>
    <body>
        <div class="container">
            <!-- 添加用户 -->
            <div class="row">
                <div class="col-md-4 col-md-offset-4">
                    <form role="form" action="/user/add" method="post">
                        <div class="form-group">
                            <label for="username">Username</label>
                            <input type="text" class="form-control" id="username" name="username" placeholder="Username" />
                        </div>
                        <div class="form-group">
                            <label for="Nickname">Username</label>
                            <input type="text" class="form-control" id="nickname" name="nickname" placeholder="Nickname" />
                        </div>
                        <div class="form-group">
                            <label for="password">Password</label>
                            <input type="password" class="form-control" id="password" name="password" placeholder="Password" />
                        </div>
                        <div class="form-group">
                            <label for="repassword">RePassword</label>
                            <input type="password" class="form-control" id="repassword" name="repassword" placeholder="RePassword" />
                        </div>
                        <button type="submit" class="btn btn-primary">Add User</button>
                    </form>
                </div>
            </div>
            <!-- 添加 path -->
            <div class="row">
                <div class="col-md-4 col-md-offset-4">
                    <form role="form" action="/path/add" method="post">
                        <div class="form-group">
                            <label for="pathname">path</label>
                            <input type="text" class="form-control" id="pathname" name="pathname" placeholder="pathname" />
                        </div>
                        <button type="submit" class="btn btn-primary">Add Path</button>
                    </form>
                </div>
            </div>
            <!-- 绑定已有的 path 和 group -->
            <div class="row">
                <div class="col-md-4 col-md-offset-4">
                    <form role="form" action="/path/bindGroupAndPath" method="post">
                        <div class="form-group">
                            <label for="pathname">path</label>
                            <input type="text" class="form-control" id="pathname" name="pathname" placeholder="pathname" />
                        </div>
                        <div class="form-group">
                            <label for="groupname">group</label>
                            <input type="text" class="form-control" id="groupname" name="groupname" placeholder="groupname" />
                        </div>
                        <button type="submit" class="btn btn-primary">Bind</button>
                    </form>
                </div>
            </div>
            <!-- 添加 group -->
            <div class="row">
                <div class="col-md-4 col-md-offset-4">
                    <form role="form" action="/group/add" method="post">
                        <div class="form-group">
                            <label for="groupname">group</label>
                            <input type="text" class="form-control" id="groupname" name="groupname" placeholder="groupname" />
                        </div>
                        <button type="submit" class="btn btn-primary">Add</button>
                    </form>
                </div>
            </div>
        </div>
    </body>
</html>
