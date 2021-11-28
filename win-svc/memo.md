windows/svc と windows/svc/mgr あたりが Windows サービス関連のよう。

* https://go.googlesource.com/sys/+/master/windows/svc/svc_test.go

C# だと ServiceController。

* https://docs.microsoft.com/ja-jp/dotnet/api/system.serviceprocess.servicecontroller?view=dotnet-plat-ext-6.0

管理者権限が必要なので、管理者権限が必要なことをどう伝えるか。

* https://stackoverflow.com/questions/31558066/how-to-ask-for-administer-privileges-on-windows-with-go
* https://github.com/mozey/run-as-admin

