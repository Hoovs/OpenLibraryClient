FROM scratch
COPY build/OpenLibraryServer /OpenLibraryServer
COPY wishList.sqlite3 /
ENTRYPOINT ["/OpenLibraryServer"]