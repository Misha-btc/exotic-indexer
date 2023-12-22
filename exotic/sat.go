package exotic

import "github.com/bitgemtech/exotic-indexer/ordinals"

const DificultyAdjustmentInterval = 2016
const HalvingInterval = 210000
const CycleInterval = HalvingInterval * 6

var PizzaRanges = ReadRangesFromOrdResponse(PIZZA_RANGES)
var NakamotoBlocks = []int64{9, 286, 688, 877, 1760, 2459, 2485, 3479, 5326, 9443, 9925, 10645, 14450, 15625, 15817, 19093, 23014, 28593, 29097}
var FirstTransactionRanges = []*ordinals.Range{
	{
		Start: 45000000000,
		Size:  1000000000,
	},
}

var HitmanRanges = SatingRangesToOrdinalsRanges(hitmanSatingRanges)
var JpegRanges = SatingRangesToOrdinalsRanges(jpegSatingRanges)

type Sat struct {
    ID              int64 // Уникальный идентификатор сатоши
    TransactionCount int  // Количество транзакций, в которых участвовал сатоши
}

func (s *Sat) Epoch() Epoch {
    return EpochFromSat(s.ID)
}

func (s *Sat) EpochPosition() int64 {
    r := s.ID - s.Epoch().GetStartingSat()
    return int64(r)
}

func (s *Sat) Height() int64 {
    r := int64(s.Epoch()) * HalvingInterval
    sub := s.Epoch().GetSubsidy()
    p := s.EpochPosition() / sub
    return p + r
}

func (s *Sat) IsFirstSatInBlock() bool {
    sub := s.Epoch().GetSubsidy()
    return int64(s.ID)%sub == 0
}

func (s *Sat) IncrementTransactionCount() {
    s.TransactionCount++
}

func ProcessTransaction(tx Transaction) {
    for _, output := range tx.Outputs {
        // Предполагается, что вы можете получить объект Sat из output
        satoshi := GetSatoshiFromOutput(output)
        satoshi.IncrementTransactionCount()
        // Обновите сатоши в вашей базе данных или кэше
    }
}

