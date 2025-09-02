{
  lib,
  buildGoModule,
  color ? "6",
  enableUpdateCmd ? false,
}:
buildGoModule rec {
  pname = "gotcha" + lib.optionalString enableUpdateCmd "-update";
  version = "2.3.1";
  src = ./..;
  vendorHash = "sha256-hpAsYPhiYnTpY5Z7QZz9cr5RtleHnR1ezgoVaQ+cvp0=";

  env.CGO_ENABLED = 0;

  tags = lib.optional enableUpdateCmd [
    "update"
  ];

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.color=${color}'"
    "-X 'github.com/MrSom3body/gotcha/cmd.version=v${version}'"
    "-extldflags '-static'"
  ];

  installPhase = ''
    runHook preInstall

    install -Dm755 $GOPATH/bin/gotcha $out/bin/${pname}

    runHook postInstall
  '';

  postInstall = ''
    mkdir -p $out/share/bash-completion/completions
    mkdir -p $out/share/zsh/site-functions
    mkdir -p $out/share/fish/vendor_completions.d

    $out/bin/${pname} completion bash > $out/share/bash-completion/completions/${pname}
    $out/bin/${pname} completion zsh > $out/share/zsh/site-functions/_${pname}
    $out/bin/${pname} completion fish > $out/share/fish/vendor_completions.d/${pname}.fish
  '';

  meta.mainProgram = pname;
}
