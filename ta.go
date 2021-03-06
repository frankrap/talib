package talib

// TA-LIB api calls

//sys ta_Initialize() (retCode int, err error) = talib.Z_Initialize
//sys ta_Shutdown() (retCode int, err error) = talib.Z_Shutdown
//sys ta_Acos(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Acos
//sys ta_Ad(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ad
//sys ta_Add(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Add
//sys ta_AdOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, optInFastPeriod int, optInSlowPeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AdOsc
//sys ta_Adx(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Adx
//sys ta_Adxr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Adxr
//sys ta_Apo(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Apo
//sys ta_AroOn(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outAroonDown *float64, outAroonUp *float64) (retCode int, err error) = talib.Z_AroOn
//sys ta_AroOnOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AroOnOsc
//sys ta_ASin(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_ASin
//sys ta_Atan(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Atan
//sys ta_Atr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Atr
//sys ta_AvgPrice(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AvgPrice
//sys ta_BBands(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDevUp float64, optInNbDevDn float64, optInMAType int, outBegIdx *int, outNBElement *int, outRealUpperBand *float64, outRealMiddleBand *float64, outRealLowerBand *float64) (retCode int, err error) = talib.Z_BBands
//sys ta_Beta(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Beta
//sys ta_Bop(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Bop
//sys ta_Cci(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cci
//sys ta_Cdl2Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl2Crows
//sys ta_Cdl3BlackCrows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3BlackCrows
//sys ta_Cdl3InSide(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3InSide
//sys ta_Cdl3LineStrike(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3LineStrike
//sys ta_Cdl3OutSide(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3OutSide
//sys ta_Cdl3StarsInSouth(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3StarsInSouth
//sys ta_Cdl3WhiteSoldiers(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3WhiteSoldiers
//sys ta_CdlAbandOnedBaBy(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlAbandOnedBaBy
//sys ta_CdlAdvanceBlock(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlAdvanceBlock
//sys ta_CdlBeltHold(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlBeltHold
//sys ta_CdlBreakaway(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlBreakaway
//sys ta_CdlCloSingMarubozu(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlCloSingMarubozu
//sys ta_CdlCOncealBaBySwall(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlCOncealBaBySwall
//sys ta_CdlCounterattack(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlCounterattack
//sys ta_CdlDarkCloudCover(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDarkCloudCover
//sys ta_CdlDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDoji
//sys ta_CdlDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDojiStar
//sys ta_CdlDragOnflyDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDragOnflyDoji
//sys ta_CdlEngulfing(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlEngulfing
//sys ta_CdlEveningDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlEveningDojiStar
//sys ta_CdlEveningStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlEveningStar
//sys ta_CdlGapSidesideWhite(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlGapSidesideWhite
//sys ta_CdlGravestOneDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlGravestOneDoji
//sys ta_CdlHammer(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHammer
//sys ta_CdlHangingMan(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHangingMan
//sys ta_CdlHarami(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHarami
//sys ta_CdlHaramiCross(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHaramiCross
//sys ta_CdlHighWave(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHighWave
//sys ta_CdlHikkake(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHikkake
//sys ta_CdlHikkakeMod(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHikkakeMod
//sys ta_CdlHoMingPigeOn(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHoMingPigeOn
//sys ta_CdlIdentical3Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlIdentical3Crows
//sys ta_CdlinNeck(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlinNeck
//sys ta_CdlinvertedHammer(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlinvertedHammer
//sys ta_CdlKicking(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlKicking
//sys ta_CdlKickingByLength(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlKickingByLength
//sys ta_CdlLadderBottom(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLadderBottom
//sys ta_CdlLOngLeggedDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLOngLeggedDoji
//sys ta_CdlLOngLine(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLOngLine
//sys ta_CdlMarubozu(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMarubozu
//sys ta_CdlMatchingLow(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMatchingLow
//sys ta_CdlMatHold(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMatHold
//sys ta_CdlMorningDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMorningDojiStar
//sys ta_CdlMorningStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMorningStar
//sys ta_CdlOnNeck(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlOnNeck
//sys ta_CdlPiercing(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlPiercing
//sys ta_CdlRickshawMan(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlRickshawMan
//sys ta_CdlRiseFall3Methods(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlRiseFall3Methods
//sys ta_CdlSeparatingLines(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlSeparatingLines
//sys ta_CdlShootingStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlShootingStar
//sys ta_CdlShortLine(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlShortLine
//sys ta_CdlSpinningTop(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlSpinningTop
//sys ta_CdlStalledPattern(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlStalledPattern
//sys ta_CdlStickSandwich(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlStickSandwich
//sys ta_CdlTakuri(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlTakuri
//sys ta_CdlTasukiGap(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlTasukiGap
//sys ta_CdlThrusting(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlThrusting
//sys ta_CdltriStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdltriStar
//sys ta_CdlUnique3River(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlUnique3River
//sys ta_CdlupSideGap2Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlupSideGap2Crows
//sys ta_CdlxSideGap3Methods(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlxSideGap3Methods
//sys ta_Ceil(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ceil
//sys ta_Cmo(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cmo
//sys ta_Correl(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Correl
//sys ta_Cos(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cos
//sys ta_Cosh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cosh
//sys ta_Dema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Dema
//sys ta_Div(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Div
//sys ta_Dx(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Dx
//sys ta_Ema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ema
//sys ta_Exp(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Exp
//sys ta_Floor(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Floor
//sys ta_HtDcPeriod(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtDcPeriod
//sys ta_HtDcPhase(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtDcPhase
//sys ta_HtPhasor(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outInPhase *float64, outQuadrature *float64) (retCode int, err error) = talib.Z_HtPhasor
//sys ta_HtSine(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outSine *float64, outLeadSine *float64) (retCode int, err error) = talib.Z_HtSine
//sys ta_HtTrendLine(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtTrendLine
//sys ta_HtTrendMode(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_HtTrendMode
//sys ta_Kama(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Kama
//sys ta_LinearReg(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LinearReg
//sys ta_LinearRegAngle(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LinearRegAngle
//sys ta_LinearRegIntercept(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LinearRegIntercept
//sys ta_LinearRegSlope(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LinearRegSlope
//sys ta_Ln(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ln
//sys ta_Log10(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Log10
//sys ta_Ma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ma
//sys ta_Macd(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInSignalPeriod int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_Macd
//sys ta_MacdExt(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInFastMAType int, optInSlowPeriod int, optInSlowMAType int, optInSignalPeriod int, optInSignalMAType int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_MacdExt
//sys ta_MacdFix(startIdx int, endIdx int, inReal *float64, optInSignalPeriod int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_MacdFix
//sys ta_Mama(startIdx int, endIdx int, inReal *float64, optInFastLimit float64, optInSlowLimit float64, outBegIdx *int, outNBElement *int, outMAMA *float64, outFAMA *float64) (retCode int, err error) = talib.Z_Mama
//sys ta_Mavp(startIdx int, endIdx int, inReal *float64, inPeriods *float64, optInMinPeriod int, optInMaxPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mavp
//sys ta_Max(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Max
//sys ta_MaxIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_MaxIndex
//sys ta_MedPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MedPrice
//sys ta_Mfi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mfi
//sys ta_MidPoint(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MidPoint
//sys ta_MidPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MidPrice
//sys ta_Min(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Min
//sys ta_MinIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_MinIndex
//sys ta_MinMax(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outMin *float64, outMax *float64) (retCode int, err error) = talib.Z_MinMax
//sys ta_MinMaxIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outMinIdx *int, outMaxIdx *int) (retCode int, err error) = talib.Z_MinMaxIndex
//sys ta_MinusDi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MinusDi
//sys ta_MinusDm(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MinusDm
//sys ta_Mom(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mom
//sys ta_Mult(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mult
//sys ta_Natr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Natr
//sys ta_Obv(startIdx int, endIdx int, inReal *float64, inVolume *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Obv
//sys ta_PlusDi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_PlusDi
//sys ta_PlusDm(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_PlusDm
//sys ta_Ppo(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ppo
//sys ta_Roc(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Roc
//sys ta_Rocp(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocp
//sys ta_Rocr(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocr
//sys ta_Rocr100(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocr100
//sys ta_Rsi(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rsi
//sys ta_Sar(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInAcceleration float64, optInMaximum float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sar
//sys ta_SarExt(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInStartValue float64, optInOffsetOnReverse float64, optInAccelerationInitLong float64, optInAccelerationLong float64, optInAccelerationMaxLong float64, optInAccelerationInitShort float64, optInAccelerationShort float64, optInAccelerationMaxShort float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_SarExt
//sys ta_Sin(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sin
//sys ta_Sinh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sinh
//sys ta_Sma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sma
//sys ta_Sqrt(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sqrt
//sys ta_StdDev(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDev float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_StdDev
//sys ta_Stoch(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInFastK_Period int, optInSlowK_Period int, optInSlowK_MAType int, optInSlowD_Period int, optInSlowD_MAType int, outBegIdx *int, outNBElement *int, outSlowK *float64, outSlowD *float64) (retCode int, err error) = talib.Z_Stoch
//sys ta_Stochf(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int, outBegIdx *int, outNBElement *int, outFastK *float64, outFastD *float64) (retCode int, err error) = talib.Z_Stochf
//sys ta_StochRsi(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int, outBegIdx *int, outNBElement *int, outFastK *float64, outFastD *float64) (retCode int, err error) = talib.Z_StochRsi
//sys ta_Sub(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sub
//sys ta_Sum(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sum
//sys ta_T3(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInVFactor float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_T3
//sys ta_Tan(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tan
//sys ta_Tanh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tanh
//sys ta_Tema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tema
//sys ta_Trange(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trange
//sys ta_Trima(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trima
//sys ta_Trix(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trix
//sys ta_Tsf(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tsf
//sys ta_TypPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_TypPrice
//sys ta_UltOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod1 int, optInTimePeriod2 int, optInTimePeriod3 int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_UltOsc
//sys ta_Var(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDev float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Var
//sys ta_WclPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_WclPrice
//sys ta_Willr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Willr
//sys ta_Wma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Wma
