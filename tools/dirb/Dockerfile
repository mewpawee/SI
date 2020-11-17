FROM bitnami/minideb:latest
RUN apt-get update && apt-get install -y dirb
CMD dirb "${target}" "${wordlist}" -o "${output}"
