package response

import (
    "io"
)

type Response interface{
    Parse(reader io.Reader) (response Response)
}
