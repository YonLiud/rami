import sys
import os
from PyInstaller.utils.hooks import collect_all

block_cipher = None

platform = os.environ.get('PLATFORM', 'win32')

if platform == "win32":
    exe_name = "rami.exe"
    console = True
    icon_file = "icon.ico"
elif platform == "linux":
    #TODO Finish linux build, currently not working
    exe_name = "rami"
    console = True
    icon_file = None
else:
    raise ValueError("Unknown platform specified. Use 'win32' or 'linux'.")

datas = [
    ('templates', 'templates'),
    ('routes', 'routes'),
    ('services', 'services'),
    ('static', 'static'),
]

a = Analysis(
    ['app.py'],
    pathex=[],
    binaries=[],
    datas=datas,
    hiddenimports=[],
    hookspath=[],
    runtime_hooks=[],
    excludes=[],
    win_no_prefer_redirects=False,
    win_private_assemblies=False,
    cipher=block_cipher,
)

pyz = PYZ(a.pure, a.zipped_data, cipher=block_cipher)

exe = EXE(
    pyz,
    a.scripts,
    [],
    exclude_binaries=True,
    name=exe_name,
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    upx_exclude=[],
    runtime_tmpdir=None,
    console=console,
    icon=icon_file,
)

coll = COLLECT(
    exe,
    a.binaries,
    a.zipfiles,
    a.datas,
    strip=False,
    upx=True,
    upx_exclude=[],
    name=exe_name,
)
