package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day21 extends BaseTest {

	int player1Pos = 1;
	int player2Pos = 10;

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day21/input.txt");
		final int round = 0;
		int current = 1;
		int score1 = 0;
		int score2 = 0;
		while (true) {
			final int roll1 = current * 3 + 3;
			current = current + 3;
			this.player1Pos = (this.player1Pos + roll1) % 10;
			if (this.player1Pos == 0) {
				this.player1Pos = 10;
			}
			score1 = this.player1Pos + score1;
			if (score1 >= 1000) {
				System.out.println(score2 * (current - 1));
				break;
			}
			final int roll2 = current * 3 + 3;
			current = current + 3;
			this.player2Pos = (this.player2Pos + roll2) % 10;
			if (this.player2Pos == 0) {
				this.player2Pos = 10;
			}
			score2 = this.player2Pos + score2;
			if (score2 >= 1000) {
				System.out.println(score1 * (current - 1));
				break;
			}
		}
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day21/input_sample.txt");
		final Map<String, GameState> DP = new HashMap<>();
		final GameState gs = doRound(0, 0, this.player1Pos, this.player2Pos, DP, true);
		System.out.println(gs.getMax());
	}

	public GameState doRound(final int score1, final int score2, final int player1Pos, final int player2Pos,
	                         final Map<String, GameState> DP, final boolean currentFirst) {
		if (score1 >= GameState.WIN_COUNT) {
			return new GameState(1, 0);
		}
		if (score2 >= GameState.WIN_COUNT) {
			return new GameState(0, 1);
		}
		final String key = score1 + "," + score2 + "," + player1Pos + "," + player2Pos + "," + currentFirst;
		if (DP.containsKey(key)) {
			return DP.get(key);
		}
		final GameState gs = new GameState(0, 0);
		for (int r1 = 1; r1 <= 3; r1++) {
			for (int r2 = 1; r2 <= 3; r2++) {
				for (int r3 = 1; r3 <= 3; r3++) {
					int newPlayer1Pos = (player1Pos + r1 + r2 + r3) % 10;
					if (newPlayer1Pos == 0) {
						newPlayer1Pos = 10;
					}
					final int newScore1 = score1 + newPlayer1Pos;
					final GameState gsOpp = doRound(score2, newScore1, player2Pos, newPlayer1Pos, DP, !currentFirst);
					gs.mergeState(gsOpp);
				}
			}
		}
		DP.put(key, gs);
		return gs;
	}

	@Data static class GameState {

		public static final int WIN_COUNT = 21;

		long win1;
		long win2;

		public GameState(final long win1, final long win2) {
			this.win1 = win1;
			this.win2 = win2;
		}

		public void mergeState(final GameState other) {
			this.win1 = this.win1 + other.getWin2();
			this.win2 = this.win2 + other.getWin1();
		}

		long getMax() {
			return Math.max(this.win1, this.win2);
		}

		@Override public String toString() {
			return "(" + this.win1 + ", " + this.win2 + ')';
		}
	}
}
