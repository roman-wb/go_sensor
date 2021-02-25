## TODO

- Generator by templates.xml

  - read xml
  - by xml build struct and decoder
  - add annotation and ID

- Reader PCAP by dump.pcap
- UDP reader by configuration.xml
- Storage in-memory
- Static Panel
- WS Panel (register clients)
- Stream snapshot by 100 ms

## Notes

Table comparators Templates (u)int8, (u)int16, (u)int32, (u)int64 = Go int(64) Templates decimal = Go float64 Templates string = Go string Templates bytevector = Go slice byte []byte

## Sample Read by templates

    // reader
    // handle, err := pcap.OpenOffline("sample_19.cap")
    // if err != nil {
    // 	log.Fatal(err)
    // }
    // defer handle.Close()

    // packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    // for packet := range packetSource.Packets() {
    // 	if packet.Layers()[2].LayerType() != layers.LayerTypeUDP {
    // 		continue
    // 	}

    // 	payload := packet.ApplicationLayer().Payload()
    // 	MsgSeqNum := binary.LittleEndian.Uint32(payload[:4])
    // 	PMAP := payload[4:5]
    // 	TemplateID := int(payload[5] & caster.STOP_BIT_RESET)
    // 	fmt.Printf("Header:\n  MsgSeqNum = %d PMAP = %08b TemplateID = %d\n", MsgSeqNum, PMAP, TemplateID)

    // 	tpl19 := generator.Template19{}
    // 	tpl19.Decode(payload, 6)
    // 	fmt.Printf("  %s\n", tpl19.DumpJSON())
    // }
