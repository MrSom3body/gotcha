{
  buildGoModule,
  color ? "6",
}:
buildGoModule rec {
  pname = "gotcha";
  version = "2.2.2";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  env.CGO_ENABLED = 0;

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.color=${color}'"
    "-X 'github.com/MrSom3body/gotcha/cmd.version=v${version}'"
    "-extldflags '-static'"
  ];

  postInstall = ''
    mkdir -p $out/share/bash-completion/completions
    mkdir -p $out/share/zsh/site-functions
    mkdir -p $out/share/fish/vendor_completions.d

    $out/bin/gotcha completion bash > $out/share/bash-completion/completions/gotcha
    $out/bin/gotcha completion zsh > $out/share/zsh/site-functions/_gotcha
    $out/bin/gotcha completion fish > $out/share/fish/vendor_completions.d/gotcha.fish
  '';

  meta.mainProgram = "gotcha";
}
