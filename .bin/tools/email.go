package main

import (
    "bytes"
    "net/smtp"
)

func main() {
    // Connect to the remote SMTP server.
    c, err := smtp.Dial("smtp.gmail.com:25")
    if err != nil {
      panic(err)
    }
    defer c.Close()
    // Set the sender and recipient.
    c.Mail("dajour.nash01@gmail.com")
    c.Rcpt("dajour.nash01@gmail.com")
    // Send the email body.
    wc, err := c.Data()
    if err != nil {
      panic(err)
    }
    defer wc.Close()
    buf := bytes.NewBufferString("This is the email body.")
    if _, err = buf.WriteTo(wc); err != nil {
      panic(err)
    }
}
