echo off & color 0A
:: ��Ŀ����
set PROJECT=openssl
:: �汾��ǩ github�Ͽɲ� :https://github.com/openssl/openssl/releases
set VESION=OpenSSL_1_1_0-stable
:: ��Ŀ·��
set PROJECT_PATH=%cd%
:: ������·��
set CODE_PATH="%PROJECT_PATH%\%PROJECT%_%VESION%"
:: github openssl ��Ŀ��ַ
set OPENSSL_GIT_URL=https://github.com/openssl/openssl.git
::��װ·��
set OPENSSL_INSTALL_DIR=%cd%

::��github�ϰ���ָ���汾��ȡԴ��
::��Ҫ�Ѱ�װgit
if not exist "%CODE_PATH%" (
git clone -b %VESION% https://github.com/openssl/openssl.git %CODE_PATH%
)

cd /d "%CODE_PATH%"

::ͨ��perl�ű�������������makefile
::��Ҫ�Ѱ�װ��perl��strawberry��
perl Configure VC-WIN32 --prefix=%OPENSSL_INSTALL_DIR% no-asm

:: ����VS���߼�Ŀ¼,ȡ���ڵ�����VS��װ·��
set VS_DEV_CMD="D:\Program Files (x86)\Microsoft Visual Studio 14.0\Common7\Tools\VsDevCmd.bat"
call %VS_DEV_CMD%
:: ����
nmake -f makefile
:: ����(��ѡ)
nmake test
:: ��װ
nmake install

pause