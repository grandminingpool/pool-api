// Code generated by ogen, DO NOT EDIT.

package apiModels

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Ref: #/components/schemas/Blockchain
type Blockchain struct {
	Coin       string `json:"coin"`
	Name       string `json:"name"`
	Ticker     string `json:"ticker"`
	AtomicUnit int    `json:"atomic_unit"`
}

// GetCoin returns the value of Coin.
func (s *Blockchain) GetCoin() string {
	return s.Coin
}

// GetName returns the value of Name.
func (s *Blockchain) GetName() string {
	return s.Name
}

// GetTicker returns the value of Ticker.
func (s *Blockchain) GetTicker() string {
	return s.Ticker
}

// GetAtomicUnit returns the value of AtomicUnit.
func (s *Blockchain) GetAtomicUnit() int {
	return s.AtomicUnit
}

// SetCoin sets the value of Coin.
func (s *Blockchain) SetCoin(val string) {
	s.Coin = val
}

// SetName sets the value of Name.
func (s *Blockchain) SetName(val string) {
	s.Name = val
}

// SetTicker sets the value of Ticker.
func (s *Blockchain) SetTicker(val string) {
	s.Ticker = val
}

// SetAtomicUnit sets the value of AtomicUnit.
func (s *Blockchain) SetAtomicUnit(val int) {
	s.AtomicUnit = val
}

// Merged schema.
// Ref: #/components/schemas/BlockchainCoinPrice
type BlockchainCoinPrice struct {
	Price                    float64       `json:"price"`
	PriceChange24hPercentage float64       `json:"price_change_24h_percentage"`
	Markets                  []MarketPrice `json:"markets"`
}

// GetPrice returns the value of Price.
func (s *BlockchainCoinPrice) GetPrice() float64 {
	return s.Price
}

// GetPriceChange24hPercentage returns the value of PriceChange24hPercentage.
func (s *BlockchainCoinPrice) GetPriceChange24hPercentage() float64 {
	return s.PriceChange24hPercentage
}

// GetMarkets returns the value of Markets.
func (s *BlockchainCoinPrice) GetMarkets() []MarketPrice {
	return s.Markets
}

// SetPrice sets the value of Price.
func (s *BlockchainCoinPrice) SetPrice(val float64) {
	s.Price = val
}

// SetPriceChange24hPercentage sets the value of PriceChange24hPercentage.
func (s *BlockchainCoinPrice) SetPriceChange24hPercentage(val float64) {
	s.PriceChange24hPercentage = val
}

// SetMarkets sets the value of Markets.
func (s *BlockchainCoinPrice) SetMarkets(val []MarketPrice) {
	s.Markets = val
}

// Merged schema.
// Ref: #/components/schemas/CoinPrice
type CoinPrice struct {
	Price                    float64 `json:"price"`
	PriceChange24hPercentage float64 `json:"price_change_24h_percentage"`
	Coin                     string  `json:"coin"`
}

// GetPrice returns the value of Price.
func (s *CoinPrice) GetPrice() float64 {
	return s.Price
}

// GetPriceChange24hPercentage returns the value of PriceChange24hPercentage.
func (s *CoinPrice) GetPriceChange24hPercentage() float64 {
	return s.PriceChange24hPercentage
}

// GetCoin returns the value of Coin.
func (s *CoinPrice) GetCoin() string {
	return s.Coin
}

// SetPrice sets the value of Price.
func (s *CoinPrice) SetPrice(val float64) {
	s.Price = val
}

// SetPriceChange24hPercentage sets the value of PriceChange24hPercentage.
func (s *CoinPrice) SetPriceChange24hPercentage(val float64) {
	s.PriceChange24hPercentage = val
}

// SetCoin sets the value of Coin.
func (s *CoinPrice) SetCoin(val string) {
	s.Coin = val
}

// Merged schema.
// Ref: #/components/schemas/MarketPrice
type MarketPrice struct {
	Price                    float64 `json:"price"`
	PriceChange24hPercentage float64 `json:"price_change_24h_percentage"`
	Ticker                   string  `json:"ticker"`
}

// GetPrice returns the value of Price.
func (s *MarketPrice) GetPrice() float64 {
	return s.Price
}

// GetPriceChange24hPercentage returns the value of PriceChange24hPercentage.
func (s *MarketPrice) GetPriceChange24hPercentage() float64 {
	return s.PriceChange24hPercentage
}

// GetTicker returns the value of Ticker.
func (s *MarketPrice) GetTicker() string {
	return s.Ticker
}

// SetPrice sets the value of Price.
func (s *MarketPrice) SetPrice(val float64) {
	s.Price = val
}

// SetPriceChange24hPercentage sets the value of PriceChange24hPercentage.
func (s *MarketPrice) SetPriceChange24hPercentage(val float64) {
	s.PriceChange24hPercentage = val
}

// SetTicker sets the value of Ticker.
func (s *MarketPrice) SetTicker(val string) {
	s.Ticker = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPayoutMode returns new OptPayoutMode with value set to v.
func NewOptPayoutMode(v PayoutMode) OptPayoutMode {
	return OptPayoutMode{
		Value: v,
		Set:   true,
	}
}

// OptPayoutMode is optional PayoutMode.
type OptPayoutMode struct {
	Value PayoutMode
	Set   bool
}

// IsSet returns true if OptPayoutMode was set.
func (o OptPayoutMode) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPayoutMode) Reset() {
	var v PayoutMode
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPayoutMode) SetTo(v PayoutMode) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPayoutMode) Get() (v PayoutMode, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPayoutMode) Or(d PayoutMode) PayoutMode {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPayoutsInfo returns new OptPayoutsInfo with value set to v.
func NewOptPayoutsInfo(v PayoutsInfo) OptPayoutsInfo {
	return OptPayoutsInfo{
		Value: v,
		Set:   true,
	}
}

// OptPayoutsInfo is optional PayoutsInfo.
type OptPayoutsInfo struct {
	Value PayoutsInfo
	Set   bool
}

// IsSet returns true if OptPayoutsInfo was set.
func (o OptPayoutsInfo) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPayoutsInfo) Reset() {
	var v PayoutsInfo
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPayoutsInfo) SetTo(v PayoutsInfo) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPayoutsInfo) Get() (v PayoutsInfo, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPayoutsInfo) Or(d PayoutsInfo) PayoutsInfo {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPoolFee returns new OptPoolFee with value set to v.
func NewOptPoolFee(v PoolFee) OptPoolFee {
	return OptPoolFee{
		Value: v,
		Set:   true,
	}
}

// OptPoolFee is optional PoolFee.
type OptPoolFee struct {
	Value PoolFee
	Set   bool
}

// IsSet returns true if OptPoolFee was set.
func (o OptPoolFee) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPoolFee) Reset() {
	var v PoolFee
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPoolFee) SetTo(v PoolFee) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPoolFee) Get() (v PoolFee, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPoolFee) Or(d PoolFee) PoolFee {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/PayoutMode
type PayoutMode string

const (
	PayoutModePplns PayoutMode = "pplns"
	PayoutModeProp  PayoutMode = "prop"
)

// AllValues returns all PayoutMode values.
func (PayoutMode) AllValues() []PayoutMode {
	return []PayoutMode{
		PayoutModePplns,
		PayoutModeProp,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s PayoutMode) MarshalText() ([]byte, error) {
	switch s {
	case PayoutModePplns:
		return []byte(s), nil
	case PayoutModeProp:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *PayoutMode) UnmarshalText(data []byte) error {
	switch PayoutMode(data) {
	case PayoutModePplns:
		*s = PayoutModePplns
		return nil
	case PayoutModeProp:
		*s = PayoutModeProp
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/PayoutsInfo
type PayoutsInfo struct {
	Interval  int    `json:"interval"`
	MinPayout OptInt `json:"min_payout"`
	MaxPayout OptInt `json:"max_payout"`
}

// GetInterval returns the value of Interval.
func (s *PayoutsInfo) GetInterval() int {
	return s.Interval
}

// GetMinPayout returns the value of MinPayout.
func (s *PayoutsInfo) GetMinPayout() OptInt {
	return s.MinPayout
}

// GetMaxPayout returns the value of MaxPayout.
func (s *PayoutsInfo) GetMaxPayout() OptInt {
	return s.MaxPayout
}

// SetInterval sets the value of Interval.
func (s *PayoutsInfo) SetInterval(val int) {
	s.Interval = val
}

// SetMinPayout sets the value of MinPayout.
func (s *PayoutsInfo) SetMinPayout(val OptInt) {
	s.MinPayout = val
}

// SetMaxPayout sets the value of MaxPayout.
func (s *PayoutsInfo) SetMaxPayout(val OptInt) {
	s.MaxPayout = val
}

// Ref: #/components/schemas/Pool
type Pool struct {
	Info   PoolInfo    `json:"info"`
	Stats  PoolStats   `json:"stats"`
	Slaves []PoolSlave `json:"slaves"`
}

// GetInfo returns the value of Info.
func (s *Pool) GetInfo() PoolInfo {
	return s.Info
}

// GetStats returns the value of Stats.
func (s *Pool) GetStats() PoolStats {
	return s.Stats
}

// GetSlaves returns the value of Slaves.
func (s *Pool) GetSlaves() []PoolSlave {
	return s.Slaves
}

// SetInfo sets the value of Info.
func (s *Pool) SetInfo(val PoolInfo) {
	s.Info = val
}

// SetStats sets the value of Stats.
func (s *Pool) SetStats(val PoolStats) {
	s.Stats = val
}

// SetSlaves sets the value of Slaves.
func (s *Pool) SetSlaves(val []PoolSlave) {
	s.Slaves = val
}

// Ref: #/components/schemas/PoolFee
type PoolFee struct {
	Fee     float64    `json:"fee"`
	SoloFee OptFloat64 `json:"solo_fee"`
}

// GetFee returns the value of Fee.
func (s *PoolFee) GetFee() float64 {
	return s.Fee
}

// GetSoloFee returns the value of SoloFee.
func (s *PoolFee) GetSoloFee() OptFloat64 {
	return s.SoloFee
}

// SetFee sets the value of Fee.
func (s *PoolFee) SetFee(val float64) {
	s.Fee = val
}

// SetSoloFee sets the value of SoloFee.
func (s *PoolFee) SetSoloFee(val OptFloat64) {
	s.SoloFee = val
}

// Ref: #/components/schemas/PoolInfo
type PoolInfo struct {
	Blockchain  OptString      `json:"blockchain"`
	Host        OptString      `json:"host"`
	Algos       []string       `json:"algos"`
	PayoutMode  OptPayoutMode  `json:"payout_mode"`
	Solo        OptBool        `json:"solo"`
	Fee         OptPoolFee     `json:"fee"`
	PayoutsInfo OptPayoutsInfo `json:"payouts_info"`
	Agents      []string       `json:"agents"`
}

// GetBlockchain returns the value of Blockchain.
func (s *PoolInfo) GetBlockchain() OptString {
	return s.Blockchain
}

// GetHost returns the value of Host.
func (s *PoolInfo) GetHost() OptString {
	return s.Host
}

// GetAlgos returns the value of Algos.
func (s *PoolInfo) GetAlgos() []string {
	return s.Algos
}

// GetPayoutMode returns the value of PayoutMode.
func (s *PoolInfo) GetPayoutMode() OptPayoutMode {
	return s.PayoutMode
}

// GetSolo returns the value of Solo.
func (s *PoolInfo) GetSolo() OptBool {
	return s.Solo
}

// GetFee returns the value of Fee.
func (s *PoolInfo) GetFee() OptPoolFee {
	return s.Fee
}

// GetPayoutsInfo returns the value of PayoutsInfo.
func (s *PoolInfo) GetPayoutsInfo() OptPayoutsInfo {
	return s.PayoutsInfo
}

// GetAgents returns the value of Agents.
func (s *PoolInfo) GetAgents() []string {
	return s.Agents
}

// SetBlockchain sets the value of Blockchain.
func (s *PoolInfo) SetBlockchain(val OptString) {
	s.Blockchain = val
}

// SetHost sets the value of Host.
func (s *PoolInfo) SetHost(val OptString) {
	s.Host = val
}

// SetAlgos sets the value of Algos.
func (s *PoolInfo) SetAlgos(val []string) {
	s.Algos = val
}

// SetPayoutMode sets the value of PayoutMode.
func (s *PoolInfo) SetPayoutMode(val OptPayoutMode) {
	s.PayoutMode = val
}

// SetSolo sets the value of Solo.
func (s *PoolInfo) SetSolo(val OptBool) {
	s.Solo = val
}

// SetFee sets the value of Fee.
func (s *PoolInfo) SetFee(val OptPoolFee) {
	s.Fee = val
}

// SetPayoutsInfo sets the value of PayoutsInfo.
func (s *PoolInfo) SetPayoutsInfo(val OptPayoutsInfo) {
	s.PayoutsInfo = val
}

// SetAgents sets the value of Agents.
func (s *PoolInfo) SetAgents(val []string) {
	s.Agents = val
}

type PoolSlave jx.Raw

type PoolStats jx.Raw