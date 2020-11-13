From alpine:latest
RUN apk add --no-cache libc6-compat openssl coreutils file curl bash
COPY ./dnsmap .
COPY ./doUpload.sh .
#CMD ./dnsmap "${targetDomain}" -r ./"${resultFile}"; bash doUpload.sh "${accessKey}"   "${secretKey}" "${spaceName}" "${resultFile}" "${resultFile}"

