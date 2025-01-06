package chartsServices

import (
	"context"
	"fmt"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainService struct{}

func (s *BlockchainService) getProtoClient(blockchain *blockchains.Blockchain) chartsProto.ChartsServiceClient {
	return chartsProto.NewChartsServiceClient(blockchain.GetConnection())
}

func (s *BlockchainService) getChartPeriodProto(period *apiModels.ChartPeriod) *chartsProto.ChartPeriod {
	switch *period {
	case apiModels.ChartPeriodHour:
		return chartsProto.ChartPeriod_Hour.Enum()
	case apiModels.ChartPeriodWeek:
		return chartsProto.ChartPeriod_Week.Enum()
	case apiModels.ChartPeriodMonth:
		return chartsProto.ChartPeriod_Day.Enum()
	default:
		return chartsProto.ChartPeriod_Day.Enum()
	}
}

func (s *BlockchainService) GetPoolStatsChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	solo *apiModels.OptBool,
) ([]*chartsProto.PoolStatsPoint, error) {
	client := s.getProtoClient(blockchain)
	poolStatsPoints, err := client.GetPoolStats(ctx, &chartsProto.GetPoolStatsRequest{
		Period: *s.getChartPeriodProto(period),
		Solo:   solo.Value,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool stats chart (blockchain: %s), error: %w", blockchain.GetInfo().Blockchain, err)
	}

	return poolStatsPoints.Points, nil
}

func (s *BlockchainService) GetRoundsChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
) ([]*chartsProto.RoundsPoint, error) {
	client := s.getProtoClient(blockchain)
	roundsPoints, err := client.GetRounds(ctx, &chartsProto.GetRoundsRequest{
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get rounds chart (blockchain: %s), error: %w", blockchain.GetInfo().Blockchain, err)
	}

	return roundsPoints.Points, nil
}

func (s *BlockchainService) GetMinerHashratesChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner string,
	solo *apiModels.OptBool,
) ([]*chartsProto.MinerHashratesPoint, error) {
	client := s.getProtoClient(blockchain)
	minerHashratesPoints, err := client.GetMinerHashrates(ctx, &chartsProto.GetMinerChartRequest{
		Miner:  miner,
		Solo:   solo.Value,
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get miner (address: %s) hashrates chart (blockchain: %s), error: %w", miner, blockchain.GetInfo().Blockchain, err)
	}

	return minerHashratesPoints.Points, nil
}

func (s *BlockchainService) GetMinerWorkerHashratesChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner, worker string,
) ([]*chartsProto.MinerHashratesPoint, error) {
	client := s.getProtoClient(blockchain)
	minerWorkerHashratesPoints, err := client.GetMinerWorkerHashrates(ctx, &chartsProto.GetMinerWorkerChartRequest{
		Miner:  miner,
		Worker: worker,
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get miner (address: %s, worker: %s) hashrates chart (blockchain: %s), error: %w", miner, worker, blockchain.GetInfo().Blockchain, err)
	}

	return minerWorkerHashratesPoints.Points, nil
}

func (s *BlockchainService) GetMinerSharesChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner string,
	solo *apiModels.OptBool,
) ([]*chartsProto.MinerSharesPoint, error) {
	client := s.getProtoClient(blockchain)
	minerSharesPoints, err := client.GetMinerShares(ctx, &chartsProto.GetMinerChartRequest{
		Miner:  miner,
		Solo:   solo.Value,
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get miner (address: %s) shares chart (blockchain: %s), error: %w", miner, blockchain.GetInfo().Blockchain, err)
	}

	return minerSharesPoints.Points, nil
}

func (s *BlockchainService) GetMinerWorkerSharesChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner, worker string,
) ([]*chartsProto.MinerSharesPoint, error) {
	client := s.getProtoClient(blockchain)
	minerWorkerSharesPoints, err := client.GetMinerWorkerShares(ctx, &chartsProto.GetMinerWorkerChartRequest{
		Miner:  miner,
		Worker: worker,
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get miner (address: %s, worker: %s) shares chart (blockchain: %s), error: %w", miner, worker, blockchain.GetInfo().Blockchain, err)
	}

	return minerWorkerSharesPoints.Points, nil
}

func (s *BlockchainService) GetMinerProfitabilitiesChartPoints(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner string,
	solo bool,
) ([]*chartsProto.MinerProfitabilityPoint, error) {
	client := s.getProtoClient(blockchain)
	minerProfitabilitiesPoints, err := client.GetMinerProfitabilities(ctx, &chartsProto.GetMinerChartRequest{
		Miner:  miner,
		Solo:   solo,
		Period: *s.getChartPeriodProto(period),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get miner (address: %s) profitabilities chart (blockchain: %s), error: %w", miner, blockchain.GetInfo().Blockchain, err)
	}

	return minerProfitabilitiesPoints.Points, nil
}
