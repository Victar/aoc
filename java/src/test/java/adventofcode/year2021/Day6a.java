package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day6a extends BaseTest {

	@Test public void runSolver() throws Exception {
		System.out.println(countFishes(80));
	}

	@Test public void runGold() throws Exception {
		System.out.println(countFishes(256));
	}

	public long countFishes(final int round) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day6/input.txt");
		final String[] arr = data.get(0).split(",");
		final Fishes fishes = new Fishes();
		for (final String input : arr) {
			fishes.addFish(Integer.parseInt(input));
		}
		for (int i = 0; i < round; i++) {
			fishes.doRound();
		}
		return fishes.total();
	}

	@Data public static class Fishes {

		long[] ages = new long[9];

		public void addFish(final int age) {
			this.ages[age] = this.ages[age] + 1;
		}

		public void doRound() {
			final long age0init = this.ages[0];
			for (int i = 0; i < this.ages.length - 1; i++) {
				this.ages[i] = this.ages[i + 1];
			}
			this.ages[6] = this.ages[6] + age0init;
			this.ages[8] = age0init;
		}

		public long total() {
			long total = 0;
			for (final long current : this.ages) {
				total += current;
			}
			return total;
		}

	}

}
