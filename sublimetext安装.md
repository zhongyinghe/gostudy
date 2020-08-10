sublime text的安装
-----
##### 1、安装sublime text

##### 2、安装Package Control（sublime的插件管理）
```
使用 ctrl+`，调起控制台
```
```
#安装代码：
import urllib.request,os,hashlib; h = '6f4c264a24d933ce70df5dedcf1dcaee' + 'ebe013ee18cced0ef93d5f746d80ef60'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://packagecontrol.io/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)
```
```
回车确认后开始安装
运行完后，在菜单preferences下会出现Package Control
输入install，选中install Package回车
在运行后会调起安装插件的窗口
```

##### 3、安装gosublime
```
步骤一:进入sublime的package目录，可以通过菜单Preferences=》Browse Package，快速跳转

步骤二:使用git bash，cd到Package目录,
运行 git clone https://margo.sh/GoSublime ，将代码下载到package目录;
下载完后，sublime会自动检测到插件并进行安装。如果没有安装go语言环境，会报错

步骤三:打开默认配置文件进行修改
"gscomplete_enabled": false,
// Whether or not gsfmt is enabled
"fmt_enabled": false,
改为:
"gscomplete_enabled": true,
 // Whether or not gsfmt is enabled
"fmt_enabled": true,

步骤四: 建立margo.go文件
通过preference->browse package进入package目录,接着进入GoSublime->src->新建一个margo文件夹，
把margo.sh->extension-example文件夹里的extension-example.go复制到margo文件夹里，并命名为margo.go;
退出再进去sublime

步骤五: 配置环境变量
点击preference->package setting->gosublime->setting user
添加
{
    "env": {
        "GOPATH": "//在环境变量里查看添加进去哦",
        "GOROOT": "//例如E:/SDK"
    }
}
```
