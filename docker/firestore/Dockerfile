FROM google/cloud-sdk:alpine
RUN apk update
RUN apk --no-cache add openjdk8
RUN gcloud components update
RUN gcloud components install beta cloud-firestore-emulator --quiet
COPY start-firestore.sh .
ENTRYPOINT [ "./start-firestore.sh" ]
