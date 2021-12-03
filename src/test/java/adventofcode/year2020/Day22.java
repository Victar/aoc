package adventofcode.year2020;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day22 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day22/input.txt");
		final List<Integer> player1List = new ArrayList<>();
		final List<Integer> player2List = new ArrayList<>();

		List<Integer> currentData = new ArrayList<>();
		for (final String input : data) {
			if (input.equalsIgnoreCase("Player 2:")) {
				player1List.addAll(currentData);
				currentData = new ArrayList<>();
			}
			try {
				final int currentNumber = Integer.valueOf(input);
				currentData.add(currentNumber);
			} catch (final Exception e) {
			}
		}
		player2List.addAll(currentData);
		System.out.println(player1List);
		System.out.println(player2List);
		while (player1List.size() > 0 && player2List.size() > 0) {
			doRound(player1List, player2List);
		}
		System.out.println(player1List);
		System.out.println(player2List);

		System.out.println(getPlayerScore(player1List));
		System.out.println(getPlayerScore(player2List));

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day22/input.txt");
		final ArrayList<Integer> player1List = new ArrayList<>();
		final ArrayList<Integer> player2List = new ArrayList<>();

		List<Integer> currentData = new ArrayList<>();
		for (final String input : data) {
			if (input.equalsIgnoreCase("Player 2:")) {
				player1List.addAll(currentData);
				currentData = new ArrayList<>();
			}
			try {
				final int currentNumber = Integer.valueOf(input);
				currentData.add(currentNumber);
			} catch (final Exception e) {
			}
		}
		player2List.addAll(currentData);
		final List<String> combinations = new ArrayList<>();
		doGoldRound(player1List, player2List, combinations);
		System.out.println(getPlayerScore(player1List));
		System.out.println(getPlayerScore(player2List));

	}

	public boolean doGoldRound(final List<Integer> player1List, final List<Integer> player2List, final List<String> combinations) {
		boolean isFirstWin = false;
		while (player1List.size() > 0 && player2List.size() > 0) {
			final String combination = player1List.toString() + player2List;
			if (combinations.contains(combination)) {
				return true;
			} else {
				combinations.add(combination);
			}
			final int player1Card = player1List.remove(0);
			final int player2Card = player2List.remove(0);
			if (player1List.size() >= player1Card && player2List.size() >= player2Card) {

				final ArrayList<Integer> player1Sublist = new ArrayList<>();
				for (int i = 0; i < player1Card; i++) {
					player1Sublist.add(player1List.get(i));
				}
				final ArrayList<Integer> player2Sublist = new ArrayList<>();
				for (int i = 0; i < player2Card; i++) {
					player2Sublist.add(player2List.get(i));
				}
				final List<String> combinationsSub = new ArrayList<>();
				isFirstWin = doGoldRound(player1Sublist, player2Sublist, combinationsSub);
			} else {
				isFirstWin = player1Card > player2Card;
			}
			if (isFirstWin) {
				player1List.add(player1Card);
				player1List.add(player2Card);
			} else {
				player2List.add(player2Card);
				player2List.add(player1Card);
			}
		}
		return isFirstWin;
	}

	public int getPlayerScore(final List<Integer> playerList) {
		int score = 0;
		for (int i = 0; i < playerList.size(); i++) {
			score += playerList.get(i) * (playerList.size() - i);
		}
		return score;
	}

	public void doRound(final List<Integer> player1List, final List<Integer> player2List) {
		final int player1Card = player1List.remove(0);
		final int player2Card = player2List.remove(0);
		if (player1Card > player2Card) {
			player1List.add(player1Card);
			player1List.add(player2Card);
		} else {
			player2List.add(player2Card);
			player2List.add(player1Card);
		}
	}

}

