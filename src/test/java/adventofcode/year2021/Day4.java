package adventofcode.year2021;

import java.util.ArrayList;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day4 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day4/input.txt");
		final String mask = data.get(0);
		final List<Board> boards = new ArrayList();
		for (int i = 2; i < data.size(); i = i + 6) {
			boards.add(new Board(data.get(i), data.get(i + 1), data.get(i + 2), data.get(i + 3), data.get(i + 4)));
		}
		final String[] maskArr = mask.split(",");
		for (int i = 0; i < maskArr.length; i++) {
			final int current = Integer.parseInt(maskArr[i]);
			final boolean found = doRoundSilver(boards, current);
			if (found) {
				return;
			}
		}

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day4/input.txt");
		final String mask = data.get(0);
		List<Board> boards = new ArrayList();
		for (int i = 2; i < data.size(); i = i + 6) {
			boards.add(new Board(data.get(i), data.get(i + 1), data.get(i + 2), data.get(i + 3), data.get(i + 4)));
		}
		final String[] maskArr = mask.split(",");
		for (int i = 0; i < maskArr.length; i++) {
			final int current = Integer.parseInt(maskArr[i]);
			boards = doRoundGold(boards, current);
		}

	}

	public List<Board> doRoundGold(final List<Board> boards, final int number) {
		final List<Board> filterList = new ArrayList(boards);
		for (final Board board : boards) {
			board.doRound(number);
			if (!board.isActive()) {
				filterList.remove(board);
			}
			if (filterList.isEmpty()) {
				System.out.println(board.getScore());
			}
		}
		return filterList;
	}

	public boolean doRoundSilver(final List<Board> boards, final int number) {
		for (final Board board : boards) {
			board.doRound(number);
			if (!board.isActive()) {
				System.out.println(board.getScore());
				return true;
			}
		}
		return false;
	}

	@Data public static class Cell {

		int key;
		Boolean value;

		public Cell(final int key, final boolean value) {
			this.key = key;
			this.value = value;
		}
	}

	@Data public static class Board {

		int score;
		boolean active = true;

		public static final int SIZE = 5;
		ArrayList<List<Cell>> numbers;

		public Board(final String line1, final String line2, final String line3, final String line4, final String line5) {
			this.numbers = new ArrayList();
			this.numbers.add(stringToLine(line1));
			this.numbers.add(stringToLine(line2));
			this.numbers.add(stringToLine(line3));
			this.numbers.add(stringToLine(line4));
			this.numbers.add(stringToLine(line5));
		}

		public void doRound(final Integer input) {
			for (int i = 0; i < SIZE; i++) {
				for (int j = 0; j < SIZE; j++) {
					if (this.numbers.get(i).get(j).getKey() == input) {
						if (this.active) {
							this.numbers.get(i).get(j).setValue(true);
							final int currentWinner = getWinner();
							if (currentWinner > 0) {
								this.score = currentWinner * input;
								this.active = false;
							}
						}
					}
				}
			}
		}

		public int getWinner() {
			//check raw
			for (int i = 0; i < SIZE; i++) {
				boolean win = true;
				for (int j = 0; j < SIZE; j++) {
					win = this.numbers.get(i).get(j).getValue() && win;
				}
				if (win) {
					return getWinSumm();
				}
			}
			//check column
			for (int i = 0; i < SIZE; i++) {
				boolean win = true;
				for (int j = 0; j < SIZE; j++) {
					win = this.numbers.get(j).get(i).getValue() && win;
				}
				if (win) {
					return getWinSumm();
				}
			}
			return 0;
		}

		public int getWinSumm() {
			int result = 0;
			for (int i = 0; i < SIZE; i++) {
				for (int j = 0; j < SIZE; j++) {
					if (!this.numbers.get(i).get(j).getValue()) {
						result += this.numbers.get(i).get(j).getKey();
					}
				}
			}
			return result;
		}

		private List<Cell> stringToLine(final String line) {
			final String[] numbers = line.split("\\s+");
			final List<Cell> result = new ArrayList();

			for (int i = 0; i < numbers.length; i++) {
				if (StringUtils.isNotBlank(numbers[i])) {
					result.add(new Cell(Integer.parseInt(numbers[i]), false));
				}
			}
			return result;
		}
	}

}
