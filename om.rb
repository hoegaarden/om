# This file was generated by GoReleaser. DO NOT EDIT.
class Om < Formula
  desc ""
  homepage ""
  version "7.1.1"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal-cf/om/releases/download/7.1.1/om-darwin-7.1.1.tar.gz"
    sha256 "23e5776864cc1024625b39c146c52e1e71b26a9ae6878dfcb3b776ca7700c86a"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/pivotal-cf/om/releases/download/7.1.1/om-linux-7.1.1.tar.gz"
    sha256 "3aa9364675f1fba506bf0fcd3b8fb9c45f2dc438f4fb5094a5fadaa3525f0b9b"
  end

  def install
    bin.install "om"
  end

  test do
    system "#{bin}/om --version"
  end
end
