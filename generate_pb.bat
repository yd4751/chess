ECHO off
%~dp0pb/protoc.exe --proto_path=%~dp0pb/center --go_out=%~dp0pb/ center.proto
%~dp0pb/protoc.exe --proto_path=%~dp0pb/log --go_out=%~dp0pb/ log.proto
%~dp0pb/protoc.exe --proto_path=%~dp0pb/login --go_out=%~dp0pb/ login.proto
%~dp0pb/protoc.exe --proto_path=%~dp0pb/table --go_out=%~dp0pb/ table.proto

%~dp0pb/protoc.exe --proto_path=%~dp0game_ddz/pb_client --go_out=%~dp0game_ddz/ client.proto
%~dp0pb/protoc.exe --proto_path=%~dp0game_ddz/pb_user --go_out=%~dp0game_ddz/ user.proto
REM %~dp0pb/protoc.exe --go_out=%~dp0pb %~dp0pb/log/log.proto
REM %~dp0pb/protoc.exe --go_out=%~dp0pb %~dp0pb/login/login.proto
REM %~dp0pb/protoc.exe --go_out=%~dp0pb %~dp0pb/table/table.proto

REM %~dp0pb/protoc.exe --go_out=%~dp0game_ddz %~dp0game_ddz/pb_client/client.proto
REM %~dp0pb/protoc.exe --go_out=%~dp0game_ddz %~dp0game_ddz/pb_user/user.proto

ECHO on
