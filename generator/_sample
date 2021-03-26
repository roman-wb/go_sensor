package generator

import (
	"encoding/json"
	"log"

	"github.com/roman-wb/go_sensor/caster"
)

type FieldInt struct {
	Id       int
	Value    int
	Nullable bool
}

type FieldDecimal struct {
	Id       int
	Value    float64
	Nullable bool
}

type FieldString struct {
	Id       int
	Value    string
	Nullable bool
}

type FieldByteVector struct {
	Id       int
	Value    []byte
	Nullable bool
}

type Template19 struct {
	ApplVerID    FieldString
	MessageType  FieldString
	SenderCompID FieldString
	MsgSeqNum    FieldInt
	SendingTime  FieldInt
	LastFragment FieldInt //opt
	NoMDEntries  FieldInt
	MDEntries    []Template19_MDEntries
}

type Template19_MDEntries struct {
	MDUpdateAction           FieldInt
	MDEntryType              FieldString
	SecurityID               FieldInt // opt
	SecurityIDSource         FieldInt
	Symbol                   FieldString // opt
	SecurityGroup            FieldString // opt
	ExchangeTradingSessionID FieldInt    // opt
	RptSeq                   FieldInt
	MarketDepth              FieldInt     // opt
	MDPriceLevel             FieldInt     // opt
	MDEntryID                FieldInt     // opt
	MDEntryPx                FieldDecimal // opt
	MDEntrySize              FieldInt     // opt
	MDEntryDate              FieldInt     // opt
	MDEntryTime              FieldInt
	NumberOfOrders           FieldInt     // opt
	MDEntryTradeType         FieldString  // opt
	TrdType                  FieldInt     // opt
	LastPx                   FieldDecimal // opt
	MDFlags                  FieldInt     // opt
	Currency                 FieldString  // opt
	Revision                 FieldInt     // opt
	OrderSide                FieldString  // opt
	MDEntrySyntheticSize     FieldInt     // opt
}

func (tpl *Template19) Decode(buffer []byte, offset int) {
	tpl.ApplVerID.Id = 1128
	tpl.ApplVerID.Value = "9"
	tpl.MessageType.Id = 1128
	tpl.MessageType.Value = "X"
	tpl.SenderCompID.Value = "MOEX"
	tpl.MsgSeqNum.Value, offset, tpl.MsgSeqNum.Nullable = caster.DecodeUInt(buffer, offset, false)
	tpl.SendingTime.Value, offset, tpl.SendingTime.Nullable = caster.DecodeUInt(buffer, offset, false)
	tpl.LastFragment.Value, offset, tpl.LastFragment.Nullable = caster.DecodeUInt(buffer, offset, true)
	tpl.NoMDEntries.Value, offset, tpl.NoMDEntries.Nullable = caster.DecodeUInt(buffer, offset, false)

	for i := 0; i < tpl.NoMDEntries.Value; i++ {
		MDEntry := Template19_MDEntries{}
		MDEntry.MDUpdateAction.Value, offset, MDEntry.MDUpdateAction.Nullable = caster.DecodeUInt(buffer, offset, false)
		MDEntry.MDEntryType.Value, offset, MDEntry.MDEntryType.Nullable = caster.DecodeString(buffer, offset, false)
		MDEntry.SecurityID.Value, offset, MDEntry.SecurityID.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.SecurityIDSource.Value = 8
		MDEntry.Symbol.Value, offset, MDEntry.Symbol.Nullable = caster.DecodeString(buffer, offset, true)
		MDEntry.SecurityGroup.Value, offset, MDEntry.SecurityGroup.Nullable = caster.DecodeString(buffer, offset, true)
		MDEntry.ExchangeTradingSessionID.Value, offset, MDEntry.ExchangeTradingSessionID.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.RptSeq.Value, offset, MDEntry.RptSeq.Nullable = caster.DecodeUInt(buffer, offset, false)
		MDEntry.MarketDepth.Value, offset, MDEntry.MarketDepth.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.MDPriceLevel.Value, offset, MDEntry.MDPriceLevel.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.MDEntryID.Value, offset, MDEntry.MDEntryID.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.MDEntryPx.Value, offset, MDEntry.MDEntryPx.Nullable = caster.DecodeDecimal(buffer, offset, true)
		MDEntry.MDEntrySize.Value, offset, MDEntry.MDEntrySize.Nullable = caster.DecodeInt(buffer, offset, true)
		MDEntry.MDEntryDate.Value, offset, MDEntry.MDEntryDate.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.MDEntryTime.Value, offset, MDEntry.MDEntryTime.Nullable = caster.DecodeUInt(buffer, offset, false)
		MDEntry.NumberOfOrders.Value, offset, MDEntry.NumberOfOrders.Nullable = caster.DecodeInt(buffer, offset, true)
		MDEntry.MDEntryTradeType.Value, offset, MDEntry.MDEntryTradeType.Nullable = caster.DecodeString(buffer, offset, true)
		MDEntry.TrdType.Value, offset, MDEntry.TrdType.Nullable = caster.DecodeInt(buffer, offset, true)
		MDEntry.LastPx.Value, offset, MDEntry.LastPx.Nullable = caster.DecodeDecimal(buffer, offset, true)
		MDEntry.MDFlags.Value, offset, MDEntry.MDFlags.Nullable = caster.DecodeInt(buffer, offset, true)
		MDEntry.Currency.Value, offset, MDEntry.Currency.Nullable = caster.DecodeString(buffer, offset, true)
		MDEntry.Revision.Value, offset, MDEntry.Revision.Nullable = caster.DecodeUInt(buffer, offset, true)
		MDEntry.OrderSide.Value, offset, MDEntry.OrderSide.Nullable = caster.DecodeString(buffer, offset, true)
		MDEntry.MDEntrySyntheticSize.Value, offset, MDEntry.MDEntrySyntheticSize.Nullable = caster.DecodeUInt(buffer, offset, true)
		tpl.MDEntries = append(tpl.MDEntries, MDEntry)
	}
}

func (tpl Template19) DumpJSON() string {
	str, err := json.MarshalIndent(tpl, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(str)
}
