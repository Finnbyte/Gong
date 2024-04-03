from os import mkdir, path
from shutil import copyfile, move, rmtree
from subprocess import Popen, check_output

GAME_NAME = "gong"
TARGET_DIR="./docs"
HTML = f"""
<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
// Polyfill
if (!WebAssembly.instantiateStreaming) {{
    WebAssembly.instantiateStreaming = async (resp, importObject) => {{
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    }};
}}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("{GAME_NAME}.wasm"), go.importObject).then(result => {{
    go.run(result.instance);
}});
</script>
"""

def run_command_with_env_args(command: str, wait_until_completion: bool, env_args: dict) -> None:
    p = Popen(command, shell=True, env=env_args, cwd=".")
    if not wait_until_completion:
        return

    (_, _) = p.communicate()
    _ = p.wait()


def get_command_output(command: str) -> str:
    output = check_output([command], shell=True, encoding="utf8", cwd=".").strip()
    return output


def main():
    target_dir_path = path.abspath(TARGET_DIR)
    try:
        try:
            # Create target dir
            mkdir(target_dir_path)
            print(f"Created directory at {target_dir_path}\n")
        except:
            pass

        # Build wasm file
        go_cache_path = get_command_output("go env GOCACHE")
        go_path = get_command_output("go env GOCACHE")
        run_command_with_env_args(f"go build -o {GAME_NAME}.wasm {GAME_NAME}", True, { "GOOS": "js", "GOARCH": "wasm", "GOCACHE": go_cache_path, "GOPATH": go_path})
        move(path.abspath(f"./{GAME_NAME}.wasm"), f"{target_dir_path}/{GAME_NAME}.wasm")
        print(f"Created wasm executable to {target_dir_path}\n")
        
        # Get file which executes the wasm
        go_root_path = get_command_output("go env GOROOT")
        src = f"{go_root_path}/misc/wasm/wasm_exec.js"
        copyfile(src, f"{target_dir_path}/wasm_exec.js")
        print(f"Copied wasm_exec.js to {target_dir_path}\n")

        # Write html contents
        with open(f"{target_dir_path}/index.html", "w") as file:
            file.write(HTML)
        print(f"Wrote html data to {target_dir_path}/index.html\n")

        print("Done.")
    except Exception as e:
        print("Exception occured:", e)
        rmtree(target_dir_path)


if __name__ == "__main__":
    main()

