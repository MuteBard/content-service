package mirrorMove

type OrderBy string

const (
    NAME  OrderBy = "NAME"
    CREATEDAT OrderBy = "CREATEDAT"
    UPDATEDAT OrderBy = "UPDATEDAT"
    SECONDS OrderBy = "SECONDS"
)